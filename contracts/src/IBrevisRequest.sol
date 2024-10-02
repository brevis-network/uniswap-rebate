// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./Lib.sol";

interface IBrevisRequest {
    function fulfillRequest(
        bytes32 _proofId,
        uint64 _nonce,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address _callbackTarget
    ) external;

    function fulfillRequests(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata _callbackTargets
    ) external;
}