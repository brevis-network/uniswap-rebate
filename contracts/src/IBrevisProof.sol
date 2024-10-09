// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./Lib.sol";

interface IBrevisProof {
    function submitAggProof(
        uint64 _chainId,
        bytes32[] calldata _requestIds,
        bytes calldata _proofWithPubInputs
    ) external;

    function validateAggProofData(uint64 _chainId, Brevis.ProofData[] calldata _proofDataArray) external view;
}