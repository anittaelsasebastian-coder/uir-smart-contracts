package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) GetDepositor(ctx contractapi.TransactionContextInterface) (string, error) {
	// PB-1: implicit Ethereum slot -> explicit Fabric GetState
	value, err := ctx.GetStub().GetState("depositor")
	if err != nil {
		return "", fmt.Errorf("failed to read depositor: %v", err)
	}
	return string(value), nil
}

func (s *SmartContract) PutDepositor(ctx contractapi.TransactionContextInterface, value string) error {
	// PB-1: implicit Ethereum slot -> explicit Fabric PutState
	return ctx.GetStub().PutState("depositor", []byte(value))
}

func (s *SmartContract) GetBeneficiary(ctx contractapi.TransactionContextInterface) (string, error) {
	// PB-1: implicit Ethereum slot -> explicit Fabric GetState
	value, err := ctx.GetStub().GetState("beneficiary")
	if err != nil {
		return "", fmt.Errorf("failed to read beneficiary: %v", err)
	}
	return string(value), nil
}

func (s *SmartContract) PutBeneficiary(ctx contractapi.TransactionContextInterface, value string) error {
	// PB-1: implicit Ethereum slot -> explicit Fabric PutState
	return ctx.GetStub().PutState("beneficiary", []byte(value))
}

func (s *SmartContract) GetArbiter(ctx contractapi.TransactionContextInterface) (string, error) {
	// PB-1: implicit Ethereum slot -> explicit Fabric GetState
	value, err := ctx.GetStub().GetState("arbiter")
	if err != nil {
		return "", fmt.Errorf("failed to read arbiter: %v", err)
	}
	return string(value), nil
}

func (s *SmartContract) PutArbiter(ctx contractapi.TransactionContextInterface, value string) error {
	// PB-1: implicit Ethereum slot -> explicit Fabric PutState
	return ctx.GetStub().PutState("arbiter", []byte(value))
}

func (s *SmartContract) GetBalance(ctx contractapi.TransactionContextInterface) (string, error) {
	// PB-1: implicit Ethereum slot -> explicit Fabric GetState
	value, err := ctx.GetStub().GetState("balance")
	if err != nil {
		return "", fmt.Errorf("failed to read balance: %v", err)
	}
	return string(value), nil
}

func (s *SmartContract) PutBalance(ctx contractapi.TransactionContextInterface, value string) error {
	// PB-1: implicit Ethereum slot -> explicit Fabric PutState
	return ctx.GetStub().PutState("balance", []byte(value))
}

func (s *SmartContract) GetIsApproved(ctx contractapi.TransactionContextInterface) (string, error) {
	// PB-1: implicit Ethereum slot -> explicit Fabric GetState
	value, err := ctx.GetStub().GetState("isApproved")
	if err != nil {
		return "", fmt.Errorf("failed to read isApproved: %v", err)
	}
	return string(value), nil
}

func (s *SmartContract) PutIsApproved(ctx contractapi.TransactionContextInterface, value string) error {
	// PB-1: implicit Ethereum slot -> explicit Fabric PutState
	return ctx.GetStub().PutState("isApproved", []byte(value))
}

func (s *SmartContract) Deposit(ctx contractapi.TransactionContextInterface) error {
	// PB-2: No direct equivalent of msg.sender in Fabric.
	// SG-1: GetClientIdentity DN string comparison used to approximate identity check.
	callerID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("failed to get identity: %v", err)
	}
	_ = callerID

	// SG-2 (PB-4): balance += msg.value
	// TODO: implement via GetState/PutState with math/big.Int for uint256 overflow semantics

	// SG-3 (PB-5): Ethereum events are permanently indexed and queryable via eth_getLogs.
	// Fabric events are subscription-based and cannot be permanently queried at contract level.
	// Apply off-chain event indexer for permanent queryability.
	eventPayload, _ := json.Marshal(map[string]string{"event": "emit Deposited(msg.sender,msg.value)"})
	ctx.GetStub().SetEvent("depositEvent", eventPayload)
	return nil
}

func (s *SmartContract) Approve(ctx contractapi.TransactionContextInterface) error {
	// SG-1 (PB-2): require(msg.sender == arbiter)
	// Approximated using GetClientIdentity().GetID() string comparison against stored arbiter.
	callerID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("failed to get identity: %v", err)
	}
	arbiterBytes, err := ctx.GetStub().GetState("arbiter")
	if err != nil {
		return fmt.Errorf("failed to read arbiter: %v", err)
	}
	if callerID != string(arbiterBytes) {
		return fmt.Errorf("SG-1: caller identity does not match arbiter; access denied")
	}

	// SG-2 (PB-4): require(balance > 0)
	// TODO: implement balance check via GetState("balance") with math/big.Int

	// SG-4 (PB-3) NON-PORTABLE: payable(beneficiary).transfer(balance)
	// No direct equivalent of Ether transfer in Hyperledger Fabric (no built-in currency).
	// TODO: implement via external token chaincode or off-chain settlement via event subscription.

	// SG-3 (PB-5): Ethereum event logs are permanently indexed and queryable via eth_getLogs.
	// Fabric events are subscription-based; missed if client not subscribed at block commit time.
	// Apply off-chain event indexer for persistent queryability.
	eventPayload, _ := json.Marshal(map[string]string{"event": "emit Approved(amount)"})
	ctx.GetStub().SetEvent("approveEvent", eventPayload)
	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode: %v", err))
	}
	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("Error starting chaincode: %v", err))
	}
}
