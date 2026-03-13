pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import {GasRebate} from "../src/GasRebate.sol";
import {IRouterBeneficiaryRegistry} from "../src/IRouterBeneficiaryRegistry.sol";
import {RouterBeneficiaryRegistry} from "../src/RouterBeneficiaryRegistry.sol";

contract DeployRebate is Script {
    address internal constant BRV_REQ = 0x842B53755936e033428334390c4b6Ea3F8d8085a;

    function run() public {
        vm.startBroadcast();

        // RouterBeneficiaryRegistry registry = new RouterBeneficiaryRegistry();
        GasRebate rebate = new GasRebate(IRouterBeneficiaryRegistry(address(0x28C0d8BCb791E8582F687d9d62396bA7A92a474E)), BRV_REQ, bytes32(0));

        // console.log("RouterBeneficiaryRegistry deployed at", address(registry));
        console.log("GasRebate deployed at", address(rebate));
        console.log("Brevis request address", BRV_REQ);

        vm.stopBroadcast();
    }
}
