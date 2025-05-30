// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.26;

struct ProofData {
    bytes32 commitHash;
    bytes32 appCommitHash; // zk-program computing circuit commit hash
    bytes32 appVkHash; // zk-program computing circuit Verify Key hash
    bytes32 smtRoot;
    bytes32 dummyInputCommitment; // zk-program computing circuit dummy input commitment
}

// func and types to interact with Brevis Proof core contract
interface IBrevisProof {
    function submitProof(uint64 _chainId, bytes calldata _proofWithPubInputs)
        external
        returns (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash);

    function submitAggProof(uint64 _chainId, bytes32[] calldata _requestIds, bytes calldata _proofWithPubInputs)
        external;

    function validateAggProofData(uint64 _chainId, ProofData[] calldata _proofDataArray) external view;
}