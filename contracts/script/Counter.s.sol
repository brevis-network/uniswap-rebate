// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {GasRebate} from "../src/GasRebate.sol";

contract Deploy is Script {
    GasRebate public counter;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        counter = new GasRebate();

        vm.stopBroadcast();
    }
}
