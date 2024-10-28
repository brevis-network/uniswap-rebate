pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import {UniEthT} from "../src/UniEthTestnet.sol";

contract Deploy is Script {
    function run() public {
        vm.startBroadcast();
        UniEthT u = new UniEthT(2958161126901270);
        console.log("UniEthT contract deployed at ", address(u));
        vm.stopBroadcast();
    }
}