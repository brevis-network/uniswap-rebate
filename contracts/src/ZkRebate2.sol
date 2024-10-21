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
    function claimWithZkProof(
        address receiver, // uni will be sent to this address
        bytes calldata _proof,
        bytes calldata _appOutput
    ) external {
        // check proof
        (, bytes32 appCommitHash, bytes32 appVkHash) = brvProof.submitProof(uint64(block.chainid), _proof);
        require(appVkHash == vkHash, "mismatch vkhash");
        require(appCommitHash == keccak256(_appOutput), "invalid circuit output");
        require(_appOutput.length >= 84, "not enough app output");
        require((_appOutput.length-20) % 64 == 0, "incorrect app output");
    
        // sender is msg.sender for Swap
        address sender = address(bytes20(_appOutput[0:20]));
        require(msg.sender==sender, "mismatch msg.sender and circuit output");

        uint256 amount = 0; // total uni
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
        if(amount > 0 ) {
            uni.transfer(receiver, amount);
        }
        // emit event?
    }
}