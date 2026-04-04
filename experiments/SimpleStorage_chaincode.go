package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}


func (s *SmartContract) GetStoredValue(ctx contractapi.TransactionContextInterface) (string, error) {
    // PB-1: implicit Ethereum slot -> explicit Fabric GetState
    value, err := ctx.GetStub().GetState("storedValue")
    if err != nil {
        return "", fmt.Errorf("failed to read storedValue: %v", err)
    }
    return string(value), nil
}

func (s *SmartContract) PutStoredValue(ctx contractapi.TransactionContextInterface, value string) error {
    // PB-1: implicit Ethereum slot -> explicit Fabric PutState
    return ctx.GetStub().PutState("storedValue", []byte(value))
}

func (s *SmartContract) Store(ctx contractapi.TransactionContextInterface) error {
    // PB-1: storedValue = _value
    // TODO: translate to GetState/PutState
    return nil
}
func (s *SmartContract) Retrieve(ctx contractapi.TransactionContextInterface) error {
    // PB-1: return storedValue -> GetState
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
