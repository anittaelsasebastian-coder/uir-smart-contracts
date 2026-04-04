import json
import sys

def load_uir(filepath):
    with open(filepath, 'r') as f:
        return json.load(f)

def generate_header(contract_name):
    return f"""package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {{
    contractapi.Contract
}}
"""

def generate_state_helper(var):
    name = var['name']
    Name = name[0].upper() + name[1:]
    return f"""
func (s *SmartContract) Get{Name}(ctx contractapi.TransactionContextInterface) (string, error) {{
    // PB-1: implicit Ethereum slot -> explicit Fabric GetState
    value, err := ctx.GetStub().GetState("{name}")
    if err != nil {{
        return "", fmt.Errorf("failed to read {name}: %v", err)
    }}
    return string(value), nil
}}

func (s *SmartContract) Put{Name}(ctx contractapi.TransactionContextInterface, value string) error {{
    // PB-1: implicit Ethereum slot -> explicit Fabric PutState
    return ctx.GetStub().PutState("{name}", []byte(value))
}}
"""

def generate_function(fn, state_vars):
    name = fn['name']
    Name = name[0].upper() + name[1:]
    lines = []
    lines.append(f"func (s *SmartContract) {Name}(ctx contractapi.TransactionContextInterface) error {{")

    for stmt in fn.get('statements', []):
        kind = stmt['kind']
        raw = stmt.get('raw', '')
        gap = stmt.get('semantic_gap', '')

        if kind == 'require' and 'msg.sender' in raw:
            lines.append(f"    // PB-2: {gap}")
            lines.append(f"    callerID, err := ctx.GetClientIdentity().GetID()")
            lines.append(f"    if err != nil {{ return fmt.Errorf(\"failed to get identity: %v\", err) }}")
            lines.append(f"    _ = callerID")

        elif kind == 'emit':
            lines.append(f"    // PB-5: {gap}")
            lines.append(f"    eventPayload, _ := json.Marshal(map[string]string{{\"event\": \"{raw}\"}})")
            lines.append(f"    ctx.GetStub().SetEvent(\"{name}Event\", eventPayload)")

        elif kind == 'transfer':
            lines.append(f"    // PB-3 NON-PORTABLE: {gap}")
            lines.append(f"    // TODO: implement via external token chaincode")

        elif kind == 'assignment':
            lines.append(f"    // PB-1: {raw}")
            lines.append(f"    // TODO: translate to GetState/PutState")

        elif kind == 'return':
            lines.append(f"    // PB-1: {raw} -> GetState")

        else:
            lines.append(f"    // {kind}: {raw}")

    lines.append("    return nil")
    lines.append("}")
    return "\n".join(lines)

def generate_main():
    return """
func main() {
    chaincode, err := contractapi.NewChaincode(&SmartContract{})
    if err != nil {
        panic(fmt.Sprintf("Error creating chaincode: %v", err))
    }
    if err := chaincode.Start(); err != nil {
        panic(fmt.Sprintf("Error starting chaincode: %v", err))
    }
}
"""

def generate_chaincode(uir_path, output_path):
    uir = load_uir(uir_path)
    contract_name = uir['contract_name']
    state_vars = uir.get('state_variables', [])
    functions = uir.get('functions', [])

    code = []
    code.append(generate_header(contract_name))

    for var in state_vars:
        code.append(generate_state_helper(var))

    for fn in functions:
        code.append(generate_function(fn, state_vars))

    code.append(generate_main())

    output = "\n".join(code)
    with open(output_path, 'w') as f:
        f.write(output)

    print(f"Chaincode written to {output_path}")

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Usage: python3 generator.py <uir_file> <output_file>")
        sys.exit(1)
    generate_chaincode(sys.argv[1], sys.argv[2])
