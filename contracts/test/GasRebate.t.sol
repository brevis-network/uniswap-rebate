// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {GasRebate} from "../src/GasRebate.sol";

contract CounterTest is Test {
    GasRebate public counter;

    function setUp() public {
        counter = new GasRebate();
        counter.setNumber(0);
    }
}
