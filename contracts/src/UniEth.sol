// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

interface IChainLink {
    function latestRoundData()
        external
        view
        returns (
            uint80 roundId,
            int256 answer,
            uint256 startedAt,
            uint256 updatedAt,
            uint80 answeredInRound
        );
}

// slot0 saves uni-eth rate
contract UniEth {
    // slot0, if 1 uni = 0.003012 eth, value is 0.003012*10^18
    int256 public answer;
    // chainlink uni-eth
    address public immutable oracle = 0xD6aA3D25116d8dA79Ea0246c4826EB951872e02e;

    function update() external {
        (, answer, , , ) = IChainLink(oracle).latestRoundData();
    }
}
