// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Escrow {
    address public depositor;
    address public beneficiary;
    address public arbiter;
    uint256 public balance;
    bool public isApproved;

    event Deposited(address indexed depositor, uint256 amount);
    event Approved (uint256 amount);

    constructor(address _beneficiary, address _arbiter){
        depositor = msg.sender;
        beneficiary = _beneficiary;
        arbiter = _arbiter;
        isApproved = false;
        balance = 0;
    }

    function deposit()public payable {
        require(msg.sender == depositor, " depositor is the only preson can deposit" );
        balance += msg.value;
        emit Deposited(msg.sender, msg.value);
    }

    function approve() public {
        require(msg.sender == arbiter, " arbiter is the only person who can approve");
        require(balance >0," there is no fund to release");
        isApproved = true;
        uint256 amount = balance;
        balance = 0;
        payable(beneficiary).transfer(amount);
        emit Approved(amount);
    }

    function getBalance()public view returns (uint256){
        return balance;
    }
}