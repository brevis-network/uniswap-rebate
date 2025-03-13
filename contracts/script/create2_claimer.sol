pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import {ClaimHelp} from "../src/ClaimHelp.sol";

// 0x11223344 salt: 0xcb25f49706fb300c92a3a6f87c37e31f367ddb229dcbd0315ec3dd8432bde3c3
// 0x123456 salt: 0x2ae2f0f5bcc1e3d9cab1536e025e6cad1f968707d5f2709178bb3efe5ea7994a
contract Deploy is Script {
    function run() public {
        vm.startBroadcast();
        // cast create2 -s 0x112233 --init-code-hash 0xd5329f02fc3251fc45e35cbb8b98d9779ab9e6f7d2433ddc14e3c5980484cc75
        // found this salt
        uint256 salt = 0x088070555e930690179eea90d94bc6fe35f350e8a7dab24f24f9be165bff974c;
        address deployer = 0x4e59b44847b379578588920cA78FbF26c0B4956C;
        ClaimHelp c = new ClaimHelp{salt: bytes32(salt)}();
        console.log("ClaimHelp contract deployed at ", address(c));
        address compA = computeAddress(deployer, salt, type(ClaimHelp).creationCode);
        console.log("Computed addr:", compA);
        console.logBytes32(keccak256(type(ClaimHelp).creationCode));
        vm.stopBroadcast();
    }
}

function computeAddress(address deployer, uint256 salt, bytes memory creationCodeWithArgs)
        pure
        returns (address hookAddress)
    {
        return address(
            uint160(uint256(keccak256(abi.encodePacked(bytes1(0xFF), deployer, salt, keccak256(creationCodeWithArgs)))))
        );
    }