# Unified Intermediate Representation for Smart Contract Portability

## Overview
This repository contains the research artefact for my MSc thesis, which focuses on designing, implementing, and experimentally validating a Unified Intermediate Representation (UIR) to support partial portability of smart contracts between blockchain platforms. The work is centred on Ethereum (Solidity) and Hyperledger Fabric (Go chaincode).

The goal is to move beyond conceptual interoperability by developing a formal and structured intermediate layer that captures essential contract logic independently of a specific blockchain.

---

## Research Motivation
Smart contract platforms differ significantly in their execution models, state management, and programming paradigms. These differences limit portability and make cross-platform reuse of contract logic difficult.

Existing approaches often remain conceptual or platform-specific. This research aims to provide a technically grounded and implementation-oriented solution by:
- Analysing execution and programming differences across platforms
- Defining a structured intermediate representation
- Building a prototype transformation pipeline
- Evaluating portability using measurable indicators

---

## Research Objectives
The main objectives of this work are:

1. Analyse and compare smart contract execution models and programming paradigms of:
   - Ethereum (Solidity)
   - Hyperledger Fabric (Go chaincode)

2. Design a formal Unified Intermediate Representation (UIR) that captures:
   - Contract structure
   - State and data types
   - Control flow and logic
   - Platform-independent semantics

3. Implement a prototype transformation pipeline:
   - Solidity → UIR → Hyperledger Fabric

4. Validate the approach using representative contract patterns.

5. Experimentally evaluate portability and identify systematic limitations.

---

## System Architecture
The proposed system is organised into the following components:

- **parser**  
  Extracts structural and semantic information from Solidity smart contracts.

- **uir-spec**  
  Formal definition of the UIR, including schemas, elements, and mapping rules.

- **uir-model**  
  Core intermediate representation data structures and semantic abstractions.

- **generator**  
  Code generation for Hyperledger Fabric from the intermediate representation.

- **experiments**  
  Test contracts, evaluation scripts, and benchmarking.

---

## Transformation Workflow
1. Input Solidity smart contract.
2. Static analysis and parsing.
3. Transformation into the Unified Intermediate Representation.
4. Platform-independent analysis and validation.
5. Code generation for Hyperledger Fabric.
6. Experimental evaluation.

---

## Evaluation Strategy
The prototype is evaluated using selected representative contract types:
- Storage contract
- Payment/transfer contract
- Escrow contract

Evaluation focuses on measurable indicators of portability, including:
- Percentage of structure and logic automatically transformed
- Categories of portable and non-portable features
- Systematic incompatibility patterns
- Required manual adaptation

This ensures repeatable and analytically structured validation.

---

## Scientific Contribution
This research contributes a structured and implementation-oriented approach to smart contract portability by:
- Designing a formal intermediate representation
- Enabling systematic analysis of platform differences
- Providing empirical insights into portability limits

---

## Practical Contribution
The thesis delivers a working prototype demonstrating partial automated transformation between Ethereum and Hyperledger Fabric, offering practical insights into cross-platform contract reuse.

---

## Current Status
Active research and prototype development.

---

## Future Work
- Support for additional blockchain platforms
- Security and vulnerability analysis at the intermediate level
- Automation of semantic verification
- Extension to more complex contract patterns

---

## Author
MSc Thesis in Business Informatics.
