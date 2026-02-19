# Design of an Intermediate Representation for Smart Contract Portability

## Overview
My MSc thesis research product, which focuses on developing, putting into practice, and experimentally testing a Unified Intermediate Representation (UIR) to provide partial portability of smart contracts between blockchain platforms, is contained in this repository. The focus of the effort is Hyperledger Fabric (Go chaincode) and Ethereum (Solidity).

By creating a formal and structured intermediary layer that captures crucial contract logic independent of a particular blockchain, the objective is to go beyond conceptual interoperability.

---

## Research Motivation
Platforms for smart contracts vary greatly in terms of programming paradigms, state management, and execution methods. These variations restrict mobility and make it challenging to reuse contract logic across platforms.

Current methods are frequently conceptual or platform-specific. The goal of this study is to offer a solution that is both technically sound and implementation-focused by:
- Examining the variations in code and execution between platforms
- Establishing a structured representation in the middle
- Constructing a transformation pipeline prototype
- Assessing portability with quantifiable metrics

---

## Research Objectives
The main objectives of this work are:

1. Examine and contrast the programming paradigms and smart contract execution models of:In terms of :
   - solidity, Ethereum
   - Hyperledger Fabric (Go chaincode)

2. Create a formal Unified Intermediate Representation (UIR) that includes the following information:
   - The construction of contracts
   - Data types and states
   - Logic and control flow
   - Semantics that are independent of platforms

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
Anitta Elasa Sebastian
MSc Thesis in Business Informatics.
