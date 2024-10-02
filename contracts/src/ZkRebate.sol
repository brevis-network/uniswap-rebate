// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

import {PoolId} from "v4-core/types/PoolId.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";

import {Brevis} from "./Lib.sol";
import {IBrevisRequest} from "./IBrevisRequest.sol";

contract ZkRebate {
    IERC20 public uni;
    IBrevisRequest public brvReq;

    // per addr, per poolid, last attested blocknum, new proof must have blocknum > lastBlockNum
    mapping(address => mapping(bytes32 => uint64)) lastBlockNum;
    // eligible pools, ie. poolkey.Hooks is non-zero.
    mapping(bytes32 => bool) poolId;

    constructor(IBrevisRequest _brv, IERC20 _uni) {
        brvReq = _brv;
        uni = _uni;
    }

    // anyone can call to add a poolid w/ non-zero hook.
    // if needed, we can query uni poolmgr to ensure poolid exsits
    function addPool(PoolKey calldata key) external {
        require(address(key.hooks)!=address(0), "pool.hooks must be non-zero");
        poolId[PoolId.unwrap(key.toId())] = true;
    }

    // to be called by msg.sender for Swap (router contract in most cases)
    // 1. relay calldata to BrevisRequest.fulfillrequests
    // 2. check msg.sender matches output[0:20]
    // 3. parse output and sum amount
    // 4. uni.transfer
    function claimWithZkProof(
        address receiver, // uni will be sent to this address
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata _callbackTargets
    ) external {
        brvReq.fulfillRequests(_proofIds, _nonces, _chainId, _proof, _proofDataArray, _appCircuitOutputs, _callbackTargets);
        uint256 amount = 0;
        for (uint256 i=0;i<_appCircuitOutputs.length;i++) {
            address sender = address(bytes20(_appCircuitOutputs[i][0:20]));
            require(msg.sender==sender, "mismatch msg.sender and circuit output");
            for (uint256 idx=20;idx<_appCircuitOutputs[i].length;idx+=64) {
                bytes32 poolid = bytes32(_appCircuitOutputs[i][idx:idx+32]);
                if(poolId[poolid]) { // valid pool
                    uint64 beginBlk = uint64(bytes8(_appCircuitOutputs[i][idx+32:idx+40]));
                    uint64 endBlk = uint64(bytes8(_appCircuitOutputs[i][idx+40:idx+48]));
                    if(beginBlk>lastBlockNum[sender][poolid]) {
                        lastBlockNum[sender][poolid] = endBlk;
                        amount += uint128(bytes16(_appCircuitOutputs[i][idx+48:idx+64]));
                    }
                }
            }
        }
        uni.transfer(receiver, amount);
    }
}