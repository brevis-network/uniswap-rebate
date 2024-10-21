// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface IBrevisProof {
    function submitProof(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs
    ) external returns (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash);
}