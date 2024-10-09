// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

import {PoolId} from "v4-core/types/PoolId.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";

import {Brevis} from "./Lib.sol";
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
    function claimWithZkProofs(
        address receiver, // uni will be sent to this address
        bytes32[] calldata _proofIds,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs
    ) external {
        // check proof
        brvProof.submitAggProof(uint64(block.chainid), _proofIds, _proof);
        brvProof.validateAggProofData(uint64(block.chainid), _proofDataArray);
        // verify data and output
        for (uint256 i=0;i<_proofIds.length;i++) {
            require(_proofDataArray[i].appVkHash == vkHash, "mismatch vkhash");
            require(_proofDataArray[i].commitHash == _proofIds[i], "invalid proofId");
            require(_proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]), "invalid circuit output");
        }

        uint256 amount = 0; // total uni
        for (uint256 i=0;i<_appCircuitOutputs.length;i++) {
            // one output has addr(20), poolid(32), fromblk(8), toblk(8), uni amount(16)
            address sender = address(bytes20(_appCircuitOutputs[i][0:20]));
            require(msg.sender==sender, "mismatch msg.sender and circuit output");
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
        uni.transfer(receiver, amount);
    }
}