// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

interface IRouterBeneficiaryRegistry {
    function beneficiaryOf(uint64 chainId, address router) external view returns (address);
}
