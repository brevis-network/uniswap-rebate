// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

import {PoolId} from "v4-core/src/types/PoolId.sol";
import {PoolKey} from "v4-core/src/types/PoolKey.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";

import {BrevisApp} from "./BrevisApp.sol";

contract GasRebate is BrevisApp {
    IERC20 public uni;
    bytes32 public vkHash;

    // per addr claimable UNI token
    mapping(address => uint256) claimable;
    // per addr, per poolid, last attested blocknum, new proof must have blocknum > lastBlockNum
    mapping(address => mapping(bytes32 => uint64)) lastBlockNum;
    // eligible pools, ie. poolkey.Hooks is non-zero.
    mapping(bytes32 => bool) poolId;

    constructor(address _brevisRequest, bytes32 _vkHash, IERC20 _uni) BrevisApp(_brevisRequest) {
        vkHash = _vkHash;
        uni = _uni;
    }
    // anyone can call to add a poolid w/ non-zero hook.
    // if needed, we can query uni poolmgr to ensure poolid exsits
    function addPool(PoolKey calldata key) external {
        require(address(key.hooks)!=address(0), "pool.hooks must be non-zero");
        poolId[PoolId.unwrap(key.toId())] = true;
    }
    
    // send amount to receiver from msg.sender's claimable
    function claim(address receiver, uint256 amount) external {
        require(claimable[msg.sender]>=amount, "amount exceeds claimable");
        claimable[msg.sender]-=amount;
        // emit event?
        uni.transfer(receiver, amount);
    }

    // brevisApp interface. _appOutput is addr,[poolid(bytes32),firstblk(u64),lastblk(u64),amount(u128)]
    function handleProofResult(bytes32 _vkHash, bytes calldata _appOutput) internal override {
        require(vkHash == _vkHash, "invalid vk");
        require(_appOutput.length >= 84, "not enough app output");
        require((_appOutput.length-20) % 64 == 0, "incorrect app output");
        // sender is msg.sender for Swap
        address sender = address(bytes20(_appOutput[0:20]));
        for (uint256 idx=20;idx<_appOutput.length;idx+=64) {
            bytes32 poolid = bytes32(_appOutput[idx:idx+32]);
            if(poolId[poolid]) {
                uint64 beginBlk = uint64(bytes8(_appOutput[idx+32:idx+40]));
                uint64 endBlk = uint64(bytes8(_appOutput[idx+40:idx+48]));
                if(beginBlk>lastBlockNum[sender][poolid]) {
                    lastBlockNum[sender][poolid] = endBlk;
                    claimable[sender] += uint128(bytes16(_appOutput[idx+48:idx+64]));
                }
            }
        }
    }
}
