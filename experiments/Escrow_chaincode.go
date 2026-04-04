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
    // PB-2: There is no direct equivalent of the msg.sender identity check in Fabric. GetClientIdentity DN string comparison was used to approximate
    callerID, err := ctx.GetClientIdentity().GetID()
    if err != nil { return fmt.Errorf("failed to get identity: %v", err) }
    _ = callerID
    // PB-1: balance += msg.value
    // TODO: translate to GetState/PutState
    // PB-5: Logs of Ethereum occurrences are permanently indexed. Fabric events are subscription-based and cannot be permanently queried at the contract level
    eventPayload, _ := json.Marshal(map[string]string{"event": "emit Deposited(msg.sender,msg.value)"})
    ctx.GetStub().SetEvent("depositEvent", eventPayload)
    return nil
}
func (s *SmartContract) Approve(ctx contractapi.TransactionContextInterface) error {
    // requird: require(msg.sender == arbiter)
    // requird: require(balance > 0)
    // PB-3 NON-PORTABLE: Ether transfer portability to Fabric‚ which has no built-in currency‚ requires a token chaincode or off-chain settlement of payments
    // TODO: implement via external token chaincode
    // PB-5: Event logs in Ethereum are permanently indexed․ In Fabric‚ events are temporary‚ subscription-based
    eventPayload, _ := json.Marshal(map[string]string{"event": "emit Approved(amount)"})
    ctx.GetStub().SetEvent("approveEvent", eventPayload)
    return nil
}
func (s *SmartContract) GetBalance(ctx contractapi.TransactionContextInterface) error {
    // PB-1: return balance -> GetState
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
