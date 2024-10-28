// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

// slot0 saves uni-eth rate
contract UniEthT {
    // slot0, if 1 uni = 0.003012 eth, value is 0.003012*10^18
    int256 public answer;

    constructor(int256 _answer) {
        answer = _answer;
    }

    function update(int256 _answer) external {
        answer = _answer;
    }
}
