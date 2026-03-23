// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    address public owner;
    uint256 public storedValue;

    constructor() {
        owner = msg.sender;
        storedValue = 0;
    }

    function store(uint256 _value) public {
        require(msg.sender == owner, "Only owner can store");
        storedValue = _value;
    }

    function retrieve() public view returns (uint256) {
        return storedValue;
    }
}
