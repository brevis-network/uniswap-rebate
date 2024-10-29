pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";

import {IBrevisProof} from "../src/IBrevisProof.sol";
import {IMyRouter} from "../src/IMyRouter.sol";
import {ZkRebate} from "../src/ZkRebate.sol";

contract Deploy is Script {
    IERC20 uni = IERC20(0x0dB4ceE042705d47Ef6C0818E82776359c3A80Ca);
    IBrevisProof brv = IBrevisProof(0x70cFEb37f84a44A25BCF7Ea172faa04A30CfD9E8);
    bytes32 vkh = 0x172fe60f977b08ba9ddc5b06d3736b4e979d1baa672d8993f72823446790bc36;
    
    function run() public {
        vm.startBroadcast();
        ZkRebate z = new ZkRebate(brv, vkh, uni);
        console.log("ZkRebate contract deployed at ", address(z));
        bytes memory _proof = hex"11032deb0646e87292fc8bd0b8295e1d4c6e5ac353b6d6e8c5f84dc128175aeb1c8b47f78bc259e8e6e5908eff305938232e44ab0e175ed73d61921bb39a98461b5743987d81ca3e79cbdea374a9832e6963f709655a58341c1518d2fe9d1aa614dd9db4f8c3313603b458e72847f7cafc6f8dccdbb85625926554a15d255bb401547b06fd68bd00257afed2aaff621295ef042852cd47fb35103952b09feb4329ca7e4dd119b99dad926bde0e98a588c7b5e03fb406ed13fe2cc4a293663b5f2881d80d4416c7432e89344e0892619e769e4d781c7efee57189b08400fe0dee23a90731ff89535bd427c531437956471c7e22bfabfc1d9913a908c3b90f54b40c982ae034ad2b9ad2b28d458841ab7a034e7946f362019ad09da48780332a121a61fd406768ca9ec4d2657527b0fa9c3e85577cee93f82ab322196ed525ead128729cccd1781e8b0580081396e658e19ab255bd4ea8be65585380148b884e000b88554d3c0978cb78cdc34cb85ba182db94a746e19c1ab9eeab2b40078923dd0b36931f217a5e2e27555789e9f503eff73fe6cd513288e1563a85c90ad3ec70202fcacb13375deff1bbe82dcb64b9f72a0364099d58538aa0f081b16daf2b30af645301cf7076eed1f6476415c8c71fdfe278126040c323b591d5cd316ce69a172fe60f977b08ba9ddc5b06d3736b4e979d1baa672d8993f72823446790bc3608b425e328b38af87090907e0a79c892f484d5524b9b09a3b99d1c6b55edac41";
        bytes[] memory _appCircuitOutputs = new bytes[](1);
        _appCircuitOutputs[0] = hex"d45f72e31e19fcf75bf47d217e500745cb36263b79e80f3ea7c2e18eb952a055f34ee2a15313adeee68804c64cdf66bfc0f6240500000000006a482e00000000006a482e0000000000000000004a3b01208e451c000000000000000000000000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000";
        bytes32[] memory _proofIds;
        IMyRouter.ProofData[] memory _proofDataArray;
        IMyRouter(0xd45f72e31e19fcf75bF47d217e500745Cb36263b).claim(address(z), 0x9F6B03Cb6d8AB8239cF1045Ab28B9Df43dfCC823, _proof, _appCircuitOutputs, _proofIds, _proofDataArray);
        // z.claimWithZkProof(, _proof, _appCircuitOutputs, _proofIds, _proofDataArray);
        vm.stopBroadcast();
    }
}