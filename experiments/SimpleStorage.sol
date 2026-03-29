// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    address public owner;
    unit256 public storedValue;
    constructor(){
        owner = msg.sender;
        storedValue = 0;
    }
    function store(unit256 value) public{
        require(msg.sender == owner, "only owner can store");
        storedValue = _value;
    }
    function retrieve()publicview returns (unit256){
        return storedValue;
    }
}