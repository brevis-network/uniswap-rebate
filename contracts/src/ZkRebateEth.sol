// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

import {IBrevisProof} from "./IBrevisProof.sol";
import {Ownable} from "./Ownable.sol";

contract ZkRebate is Ownable {
    IBrevisProof public brvProof;
    bytes32 public vkHash;

    // per addr, per poolid, last attested blocknum, new proof must have blocknum > lastBlockNum
    mapping(address => mapping(bytes32 => uint64)) public lastBlockNum;

    constructor(IBrevisProof _brv, bytes32 _vkHash) {
        brvProof = _brv;
        vkHash = _vkHash;
    }

    // to be called by msg.sender for Swap (router contract in most cases)
    // 1. submit and validate proof using brevisProof
    // 2. check msg.sender matches output[0:20]
    // 3. parse output and sum amount
    // 4. eth transfer

    // same func for signle and aggregated proof. for single proof, _proofIds and _proofDataArray are empty
    function claimWithZkProof(
        address receiver, // eth will be sent to this address
        bytes calldata _proof,
        bytes[] calldata _appCircuitOutputs,
        bytes32[] calldata _proofIds,
        IBrevisProof.ProofData[] calldata _proofDataArray
    ) external {
        uint256 amount = 0; // total eth
        if (_appCircuitOutputs.length == 1 ) {
            bytes calldata _appOutput = _appCircuitOutputs[0];
            // check proof
            (, bytes32 appCommitHash, bytes32 appVkHash) = brvProof.submitProof(uint64(block.chainid), _proof);
            require(appVkHash == vkHash, "mismatch vkhash");
            require(appCommitHash == keccak256(_appOutput), "invalid circuit output");
            amount = handleOutput(_appOutput);
            if(amount > 0 ) {
                (bool sent, ) = receiver.call{value: amount}("");
                require(sent, "failed to send eth");
            }
            return;
        }
        // batch mode
        brvProof.submitAggProof(uint64(block.chainid), _proofIds, _proof);
        brvProof.validateAggProofData(uint64(block.chainid), _proofDataArray);
        // verify data and output
        for (uint256 i=0;i<_proofIds.length;i++) {
            require(_proofDataArray[i].appVkHash == vkHash, "mismatch vkhash");
            require(_proofDataArray[i].commitHash == _proofIds[i], "invalid proofId");
            require(_proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]), "invalid circuit output");
            amount += handleOutput(_appCircuitOutputs[i]);
        }
        if(amount > 0) {
            (bool sent, ) = receiver.call{value: amount}("");
            require(sent, "failed to send eth");
        }
    }

    // parse _appOutput, return total eth amount
    // one output has addr(20), [poolid(32), fromblk(8), toblk(8), eth amount(16)]
    // circuit will ensure poolid is valid ie. PoolKey.hooks is non-zero
    function handleOutput(bytes calldata _appOutput) internal returns (uint256) {
        uint256 amount = 0;
        require(_appOutput.length >= 84, "not enough app output");
        require((_appOutput.length-20) % 64 == 0, "incorrect app output");
    
        // sender is msg.sender for Swap
        address sender = address(bytes20(_appOutput[0:20]));
        require(msg.sender==sender, "mismatch msg.sender and circuit output");

        for (uint256 idx=20;idx<_appOutput.length;idx+=64) {
            bytes32 poolid = bytes32(_appOutput[idx:idx+32]);
            if(poolid == 0) {
                break; // circuit may have zero fillings due to fixed length, ends loop early to save gas
            }
            uint64 beginBlk = uint64(bytes8(_appOutput[idx+32:idx+40]));
            uint64 endBlk = uint64(bytes8(_appOutput[idx+40:idx+48]));
            if(beginBlk>lastBlockNum[sender][poolid]) {
                lastBlockNum[sender][poolid] = endBlk;
                amount += uint128(bytes16(_appOutput[68:84]));
            }
        }
        return amount;
    }

    function setvk(bytes32 _vk) external onlyOwner {
        vkHash = _vk;
    }
    // accept eth transfer
    receive() external payable {}
    fallback() external payable {}
}