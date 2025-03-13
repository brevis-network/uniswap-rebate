// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

interface IRebateClaimer {
    function rebateClaimer() external view returns (address);
}

// simple contract call router rebateClaimer and emit Claimer event
contract ClaimHelp {
    event Claimer(address indexed router, address indexed claimer);

    function claim(address router) external {
        emit Claimer(router, IRebateClaimer(router).rebateClaimer());
    }
}