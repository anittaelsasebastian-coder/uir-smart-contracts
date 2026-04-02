import re
import json

def parse_contract(filepath):
    with open(filepath,'r') as f:
        source =f.read()

    contract_name = re.search(r'contract\s+(\w+)\s*\{',source)
    if contract_name:
        contract_name = contract_name.group(1)
    else:
        contract_name = "unknown"

    state_variables = []
    for line in source.split('\n'):
        line = line.strip()
        if re.match(r'(address|unit256|bool|string)\s+public\s+(\w+)', line):
            match = re.match(r'(address|unit256|bool|string)\s+public\s+(\w+)', line)
            var_type = match.group(1)
            var_name = match.group(2)
            barrier = "PB-2" if var_type =="address" else "PB-1"
            state_variables.append({
                "name": var_name,
                "solidity_type": var_type,
                "uir_type": var_type,
                "visibility": "public",
                "portability_barrier": barrier
            })
            