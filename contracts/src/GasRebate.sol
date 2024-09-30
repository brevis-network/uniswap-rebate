// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

contract GasRebate {
    uint256 public number;

    // per addr claimable UNI token
    mapping(address => uint256) claimable;
    // per addr, per poolid, last attested blocknum, new proof must have blocknum > lastBlockNum
    mapping(address => mapping(bytes32 => uint64)) lastBlockNum;

    function setNumber(uint256 newNumber) public {
        number = newNumber;
    }

    function increment() public {
        number++;
    }
}
