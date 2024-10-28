// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

import {PoolId} from "v4-core/types/PoolId.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";

import {IBrevisProof} from "./IBrevisProof.sol";

contract ZkRebate2 {
    IERC20 public uni;
    IBrevisProof public brvProof;
    bytes32 public vkHash;

    // per addr, per poolid, last attested blocknum, new proof must have blocknum > lastBlockNum
    mapping(address => mapping(bytes32 => uint64)) lastBlockNum;
    // eligible pools, ie. poolkey.Hooks is non-zero.
    mapping(bytes32 => bool) poolId;

    constructor(IBrevisProof _brv, bytes32 _vkHash, IERC20 _uni) {
        brvProof = _brv;
        vkHash = _vkHash;
        uni = _uni;
    }

    // anyone can call to add a poolid w/ non-zero hook.
    // if needed, we can query uni poolmgr to ensure poolid exsits
    function addPool(PoolKey calldata key) external {
        require(address(key.hooks)!=address(0), "pool.hooks must be non-zero");
        poolId[PoolId.unwrap(key.toId())] = true;
    }

    // to be called by msg.sender for Swap (router contract in most cases)
    // 1. submit and validate proof using brevisProof
    // 2. check msg.sender matches output[0:20]
    // 3. parse output and sum amount
    // 4. uni.transfer

    // same func for signle and aggregated proof. for single proof, _proofIds and _proofDataArray are empty
    function claimWithZkProof(
        address receiver, // uni will be sent to this address
        bytes calldata _proof,
        bytes[] calldata _appCircuitOutputs,
        bytes32[] calldata _proofIds,
        IBrevisProof.ProofData[] calldata _proofDataArray
    ) external {
        uint256 amount = 0; // total uni
        if (_appCircuitOutputs.length == 1 ) {
            bytes calldata _appOutput = _appCircuitOutputs[0];
            // check proof
            (, bytes32 appCommitHash, bytes32 appVkHash) = brvProof.submitProof(uint64(block.chainid), _proof);
            require(appVkHash == vkHash, "mismatch vkhash");
            require(appCommitHash == keccak256(_appOutput), "invalid circuit output");
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
                if(poolId[poolid]) {
                    uint64 beginBlk = uint64(bytes8(_appOutput[idx+32:idx+40]));
                    uint64 endBlk = uint64(bytes8(_appOutput[idx+40:idx+48]));
                    if(beginBlk>lastBlockNum[sender][poolid]) {
                        lastBlockNum[sender][poolid] = endBlk;
                        amount += uint128(bytes16(_appOutput[68:84]));
                    }
                }
            }
            // emit event?
            if(amount > 0 ) {
                uni.transfer(receiver, amount);
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
        }

        for (uint256 i=0;i<_appCircuitOutputs.length;i++) {
            // one output has addr(20), [poolid(32), fromblk(8), toblk(8), uni amount(16)]
            address sender = address(bytes20(_appCircuitOutputs[i][0:20]));
            require(msg.sender==sender, "mismatch msg.sender and circuit output");
            bytes calldata _appOutput = _appCircuitOutputs[i];
            for (uint256 idx=20;idx<_appOutput.length;idx+=64) {
                bytes32 poolid = bytes32(_appCircuitOutputs[i][20:52]);
                if(poolId[poolid]) { // valid pool
                    uint64 beginBlk = uint64(bytes8(_appCircuitOutputs[i][52:60]));
                    uint64 endBlk = uint64(bytes8(_appCircuitOutputs[i][60:68]));
                    if(beginBlk>lastBlockNum[sender][poolid]) {
                        lastBlockNum[sender][poolid] = endBlk;
                        amount += uint128(bytes16(_appCircuitOutputs[i][68:84]));
                    }
                }
            }
        }
        if(amount > 0 ) {
            uni.transfer(receiver, amount);
        }
    }
}