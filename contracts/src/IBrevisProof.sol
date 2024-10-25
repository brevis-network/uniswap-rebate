// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import "./Lib.sol";
interface IBrevisProof {
    function submitProof(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs
    ) external returns (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash);
    
    function submitAggProof(
        uint64 _chainId,
        bytes32[] calldata _requestIds,
        bytes calldata _proofWithPubInputs
    ) external;

    function validateAggProofData(uint64 _chainId, Brevis.ProofData[] calldata _proofDataArray) external view;
}