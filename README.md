# UIR: Smart Contract Portability via Unified Intermediate Representation

A prototype pipeline that transforms Ethereum smart contracts (Solidity) into Hyperledger Fabric chaincode (Go), using a platform-independent intermediate representation as the translation layer.

MSc Thesis — Business Informatics, Riga Technical University
Published thesis: https://doi.org/10.5281/zenodo.20732000

---

## What This Does

Smart contracts written for Ethereum cannot run on Hyperledger Fabric.The platforms differ in programming model, state management, and execution environment. This project builds a structured middle layer — the UIR — that captures contract logic independently of either platform, then generates working Fabric chaincode from it.

**Pipeline:**
Solidity contract → Parser → UIR (JSON) → Generator → Go chaincode

Validated against three contract types: storage, payment/transfer, and escrow.

---

## Repository Structure

| Folder | What it does |
|--------|-------------|
| `parser/` | Parses Solidity source into UIR JSON |
| `uir-spec/` | Formal UIR schema and mapping rules |
| `uir-model/` | Core UIR data structures |
| `generator/` | Generates Hyperledger Fabric Go chaincode from UIR |
| `experiments/` | Test contracts, evaluation scripts, results |

---

## Requirements

- Python 3.x
- Go (for compiling generated chaincode)

---

## How to Run

```bash
# 1. Parse a Solidity contract into UIR
cd parser
python parser.py ../experiments/escrow.sol

# 2. Generate Fabric chaincode from UIR
cd ../generator
python generator.py ../experiments/escrow_uir.json

# Output: compilable Go chaincode
```

---

## Key Findings

- Automated structural transformation:- 6 of 14 Solidity feature categories (43%) fully abstractable into UIR,4 categories (29%) approximated with documented semantic gaps, 4 categories (29%) architecturally non-portable
- All 3 generated Go chaincodes compiled successfully
- Identified 5 systematic portability barriers (PB-1 to PB-5)
- Identified 5 semantic gap categories (SG-0 to SG-4)

---

## Author

Anitta Elsa Sebastian  
MSc Business Informatics, Riga Technical University  
[Zenodo DOI](https://doi.org/10.5281/zenodo.20732000)
