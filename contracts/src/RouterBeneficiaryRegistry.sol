// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "./Ownable.sol";

contract RouterBeneficiaryRegistry is Ownable {
    mapping(uint64 => mapping(address => address)) private _beneficiaries;

    event BeneficiarySet(uint64 indexed chainId, address indexed router, address indexed beneficiary);

    function setBeneficiary(uint64 chainId, address router, address beneficiary) external onlyOwner {
        require(router != address(0), "router is zero");
        require(beneficiary != address(0), "beneficiary is zero");

        _beneficiaries[chainId][router] = beneficiary;
        emit BeneficiarySet(chainId, router, beneficiary);
    }

    function beneficiaryOf(uint64 chainId, address router) external view returns (address) {
        return _beneficiaries[chainId][router];
    }
}
