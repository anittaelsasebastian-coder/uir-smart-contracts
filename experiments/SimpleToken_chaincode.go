package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}


func (s *SmartContract) GetOwner(ctx contractapi.TransactionContextInterface) (string, error) {
    // PB-1: implicit Ethereum slot -> explicit Fabric GetState
    value, err := ctx.GetStub().GetState("owner")
    if err != nil {
        return "", fmt.Errorf("failed to read owner: %v", err)
    }
    return string(value), nil
}

func (s *SmartContract) PutOwner(ctx contractapi.TransactionContextInterface, value string) error {
    // PB-1: implicit Ethereum slot -> explicit Fabric PutState
    return ctx.GetStub().PutState("owner", []byte(value))
}


func (s *SmartContract) GetTotalSupply(ctx contractapi.TransactionContextInterface) (string, error) {
    // PB-1: implicit Ethereum slot -> explicit Fabric GetState
    value, err := ctx.GetStub().GetState("totalSupply")
    if err != nil {
        return "", fmt.Errorf("failed to read totalSupply: %v", err)
    }
    return string(value), nil
}

func (s *SmartContract) PutTotalSupply(ctx contractapi.TransactionContextInterface, value string) error {
    // PB-1: implicit Ethereum slot -> explicit Fabric PutState
    return ctx.GetStub().PutState("totalSupply", []byte(value))
}


func (s *SmartContract) GetBalances(ctx contractapi.TransactionContextInterface) (string, error) {
    // PB-1: implicit Ethereum slot -> explicit Fabric GetState
    value, err := ctx.GetStub().GetState("balances")
    if err != nil {
        return "", fmt.Errorf("failed to read balances: %v", err)
    }
    return string(value), nil
}

func (s *SmartContract) PutBalances(ctx contractapi.TransactionContextInterface, value string) error {
    // PB-1: implicit Ethereum slot -> explicit Fabric PutState
    return ctx.GetStub().PutState("balances", []byte(value))
}

func (s *SmartContract) Transfer(ctx contractapi.TransactionContextInterface) error {
    // PB-2: In Fabric, mapping read needs to become an explicit GetState call with a composite key
    callerID, err := ctx.GetClientIdentity().GetID()
    if err != nil { return fmt.Errorf("failed to get identity: %v", err) }
    _ = callerID
    // require: require(_to != address(0))
    // PB-1: balances[msg.sender] -= _amount
    // TODO: translate to GetState/PutState
    // PB-1: balances[_to] += _amount)
    // TODO: translate to GetState/PutState
    // PB-5: Logs of Ethereum occurrences are permanently indexed. Fabric events are only available with a subscription
    eventPayload, _ := json.Marshal(map[string]string{"event": "emit Transfer(msg.sender, _to, _amount)"})
    ctx.GetStub().SetEvent("transferEvent", eventPayload)
    return nil
}
func (s *SmartContract) BalanceOf(ctx contractapi.TransactionContextInterface) error {
    // PB-1: return balances[_account] -> GetState
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
