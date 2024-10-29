pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {PoolKey} from "v4-core/types/PoolKey.sol";
import {PoolId} from "v4-core/types/PoolId.sol";
import {IHooks} from "v4-core/interfaces/IHooks.sol";
import {Currency} from "v4-core/types/Currency.sol";

import {ZkRebate} from "../src/ZkRebate.sol";

contract Deploy is Script {
    PoolKey public key = PoolKey({
        currency0: Currency.wrap(0x0dB4ceE042705d47Ef6C0818E82776359c3A80Ca),
        currency1: Currency.wrap(0x7A46219950d8a9bf2186549552DA35Bf6fb85b1F),
        fee: 1000,
        tickSpacing: 60,
        hooks: IHooks(0xaA5f84D2F6A2423E50b3e598F934f6626e12CAc0)
    });
    function run() public {
        console.logBytes32(PoolId.unwrap(key.toId()));
        vm.startBroadcast();
        ZkRebate z = ZkRebate(address(0xA50e5203ECa1685B4e01A8A569dd12150A8b419D));
        z.addPool(key);
        vm.stopBroadcast();
    }
}