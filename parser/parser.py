import re
import json

def parse_contract(filepath):
    with open(filepath, 'r') as f:
        source = f.read()

    contract_name = re.search(r'contract\s+(\w+)\s*\{', source)
    if contract_name:
        contract_name = contract_name.group(1)
    else:
        contract_name = "Unknown"

    state_variables = []
    for line in source.split('\n'):
        line = line.strip()
        if re.match(r'(address|uint256|bool|string)\s+public\s+(\w+)', line):
            match = re.match(r'(address|uint256|bool|string)\s+public\s+(\w+)', line)
            var_type = match.group(1)
            var_name = match.group(2)
            barrier = "PB-2" if var_type == "address" else "PB-1"
            state_variables.append({
                "name": var_name,
                "solidity_type": var_type,
                "uir_type": var_type,
                "visibility": "public",
                "portability_barrier": barrier
            })

    uir = {
        "uir_version": "1.0",
        "contract_name": contract_name,
        "source_platform": "Ethereum",
        "source_language": "Solidity",
        "state_variables": state_variables,
        "constructor": None,
        "functions": [],
        "portability_summary": {
            "barriers_present": ["PB-1"],
            "abstractable": ["PB-1"],
            "approximable": [],
            "non_portable": []
        }
    }

    output_path = filepath.replace('.sol', '_uir.json')
    with open(output_path, 'w') as f:
        json.dump(uir, f, indent=2)

    print(f"UIR written to {output_path}")
    return uir

if __name__ == "__main__":
    parse_contract("experiments/SimpleStorage.sol")
