// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "./Ownable.sol";
import {IBrevisProof} from "./IBrevisProof.sol";
import {IRouterBeneficiaryRegistry} from "./IRouterBeneficiaryRegistry.sol";

contract GasRebate is Ownable {
    uint256 private constant _NOT_ENTERED = 1;
    uint256 private constant _ENTERED = 2;

    IBrevisProof public brvProof;
    IRouterBeneficiaryRegistry public beneficiaryRegistry;
    bytes32 public vkHash; // ensure output is from expected zk circuit
    uint256 private _reentrancyStatus = _NOT_ENTERED;
    // chainId => router => blockNum, last attested block number to avoid replay
    mapping(uint64 chainId => mapping(address router => uint64 blockNum)) public lastBlockNum;

    event Claimed(
        uint64 indexed chainId,
        address indexed router,
        address indexed recipient,
        uint256 amount
    );

    constructor(IBrevisProof _brvProof, bytes32 _vkHash) {
        brvProof = _brvProof;
        vkHash = _vkHash;
    }

    modifier nonReentrant() {
        require(_reentrancyStatus != _ENTERED, "reentrant call");
        _reentrancyStatus = _ENTERED;
        _;
        _reentrancyStatus = _NOT_ENTERED;
    }

    function claimWithZkProof(
        bytes calldata _proof,
        bytes[] calldata _appCircuitOutputs,
        bytes32[] calldata _proofIds,
        IBrevisProof.ProofData[] calldata _proofDataArray
    ) external nonReentrant {
        require(vkHash != bytes32(0), "vkHash not set");
        require(_appCircuitOutputs.length > 0, "no app output");
    
        if (_appCircuitOutputs.length == 1) {
            bytes calldata _appOutput = _appCircuitOutputs[0];
            (uint64 outputChainid, address router, uint64 fromBlk, uint64 toBlk, uint256 amount) = parseAppOutput(_appOutput);
            // check proof
            (, bytes32 appCommitHash, bytes32 appVkHash) = brvProof.submitProof(outputChainid, _proof);

            require(appVkHash == vkHash, "mismatch vkHash");
            require(appCommitHash == keccak256(_appOutput), "invalid circuit output");

            sendRebate(outputChainid, router, fromBlk, toBlk, amount);
            return;
        }
        // batch mode
        require(_proofIds.length == _appCircuitOutputs.length, "invalid proofIds length");
        require(_proofDataArray.length == _appCircuitOutputs.length, "invalid proofData length");
        uint64 chainid = uint64(bytes8(_appCircuitOutputs[0][0:8]));
        brvProof.submitAggProof(chainid, _proofIds, _proof);
        brvProof.validateAggProofData(chainid, _proofDataArray);
        // verify data and output
        for (uint256 i = 0; i < _proofIds.length; i++) {
            require(_proofDataArray[i].appVkHash == vkHash, "mismatch vkhash");
            require(_proofDataArray[i].commitHash == _proofIds[i], "invalid proofId");
            require(_proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]), "invalid circuit output");

            (uint64 outputChainid, address router, uint64 fromBlk, uint64 toBlk, uint256 amount) =
                parseAppOutput(_appCircuitOutputs[i]);
            require(outputChainid == chainid, "chainid mismatch");
            sendRebate(outputChainid, router, fromBlk, toBlk, amount);
        }
    }

    function sendRebate(uint64 chainid, address router, uint64 fromBlk, uint64 toBlk, uint256 amount) internal {
        require(fromBlk <= toBlk, "invalid block range");
        require(fromBlk > lastBlockNum[chainid][router], "begin blocknum too small");

        lastBlockNum[chainid][router] = toBlk;
        address recipient = beneficiaryRegistry.beneficiaryOf(chainid, router);
        require(recipient != address(0), "beneficiary is zero");

        if (amount > 0) {
            (bool sent,) = recipient.call{value: amount}("");
            require(sent, "failed to send eth");

            emit Claimed(chainid, router, recipient, amount);
        }
    }

    // one output has chainid(8), router(20), fromBlk(8), toBlk(8), amount(16)
    function parseAppOutput(bytes calldata _appOutput) internal pure returns (
        uint64 chainid,
        address router,
        uint64 fromBlk,
        uint64 toBlk,
        uint256 amount
    ) {
        require(_appOutput.length == 60, "incorrect app output length");
        chainid = uint64(bytes8(_appOutput[0:8]));
        router = address(bytes20(_appOutput[8:28]));
        fromBlk = uint64(bytes8(_appOutput[28:36]));
        toBlk = uint64(bytes8(_appOutput[36:44]));
        amount = uint256(uint128(bytes16(_appOutput[44:60])));
    }

    function setBrvProof(IBrevisProof _brvProof) external onlyOwner {
        brvProof = _brvProof;
    }
    function setBeneficiaryRegistry(IRouterBeneficiaryRegistry _registry) external onlyOwner {
        beneficiaryRegistry = _registry;
    }
    function setVk(bytes32 _vk) external onlyOwner {
        vkHash = _vk;
    }
    // Escape hatch function to withdraw ETH from the contract
    function withdrawETH(address _recipient) external onlyOwner {
        (bool success,) = _recipient.call{value: address(this).balance}("");
        require(success, "transfer failed");
    }
    // accept eth transfer
    receive() external payable {}
    fallback() external payable {}
}
