// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "./Ownable.sol";
import {BrevisAppZkOnly} from "./BrevisAppZkOnly.sol";
import {IRouterBeneficiaryRegistry} from "./IRouterBeneficiaryRegistry.sol";

contract GasRebate is BrevisAppZkOnly, Ownable {

    IRouterBeneficiaryRegistry public immutable beneficiaryRegistry;
    bytes32 public vkHash; // ensure output is from expected zk circuit

    // chainId => router => blockNum, last attested block number to avoid replay
    mapping(uint64 chainId => mapping(address router => uint64 blockNum)) public lastBlockNum;

    event Claimed(
        uint64 indexed chainId,
        address indexed router,
        address indexed recipient,
        uint256 amount
    );

    constructor(IRouterBeneficiaryRegistry _registry, address _brvReq, bytes32 _vkHash) BrevisAppZkOnly(_brvReq) {
        beneficiaryRegistry = _registry;
        vkHash = _vkHash;
    }

    // override logic to handle circuit output data
    function handleProofResult(bytes32 _vkHash, bytes calldata _circuitOutput) internal override {
        require(vkHash == _vkHash, "invalid vk");
        (uint64 chainId, address router, uint64 fromBlk, uint64 toBlk, uint256 amount) = parseAppOutput(_circuitOutput);
        sendRebate(chainId, router, fromBlk, toBlk, amount);
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
