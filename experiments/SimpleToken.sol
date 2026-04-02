// SPDX-License-Identifier:MIT
pragma solidity ^0.8.0;

contract SimpleToken{
    address public owner;
    uint256 public totalSupply;
    mapping(address => uint256) public balances;
    
    event Transfer(address indexed from, address indexed to, uint256 amount);

    constructor(uint256 _initialSupply) {
        owner = msg.sender;
        totalSupply = _initialSupply;
        balances[msg.sender]= _initialSupply;
    }

    function transfer(address _to, uint256 _amount) public { 
        require(balances[msg.sender] >= _amount, "Insufficient balance");
        require(_to != address(0),"Invalid recipient");
        balances[msg.sender] -= _amount;
        balances[_to] += _amount;
        emit Transfer(msg.sender,_to,_amount);
    }
    function balanceOf(address _account) public view returns(uint256) {
        return balances[_account];
    }
}