// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.13;

// to be called by router contracts
interface IMyRouter {
    struct ProofData {
        bytes32 commitHash;
        bytes32 vkHash;
        bytes32 appCommitHash; // zk-program computing circuit commit hash
        bytes32 appVkHash; // zk-program computing circuit Verify Key hash
        bytes32 smtRoot;
    }
    function claim(
        address zk,
        address receiver, // uni will be sent to this address
        bytes calldata _proof,
        bytes[] calldata _appCircuitOutputs,
        bytes32[] calldata _proofIds,
        ProofData[] calldata _proofDataArray
    ) external;
}