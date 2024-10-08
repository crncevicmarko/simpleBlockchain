# Simple Blockchain Implementation in Go

This repository contains a basic implementation of a blockchain in Go. The blockchain consists of blocks that store data, a timestamp, and a reference to the previous block's hash. The main goal of this project is to demonstrate the foundational concepts of blockchain technology and how to create a simple blockchain structure in Go.

## Features

- **Block Structure**: Each block contains:
  - `TimeStamp`: The time the block was created.
  - `Data`: The information stored in the block.
  - `PrevBlockHash`: A hash of the previous block, ensuring the integrity of the chain.
  - `Hash`: The hash of the current block, calculated based on its contents.
  - `Nonce`: A number used in the proof-of-work algorithm to find a valid hash.

- **Blockchain Structure**:
  - A blockchain consists of multiple blocks linked together.
  - The first block is known as the Genesis block.
  - Each block is mined using a proof-of-work mechanism.

- **Proof of Work**:
  - Implements a proof-of-work algorithm to secure the blockchain.
  - Uses the SHA-256 hashing algorithm to generate a hash for each block.
  - The hash is calculated from the block's timestamp, data, previous block's hash, and nonce.

- **Validation**:
  - Each block can be validated to ensure its hash is less than the target, verifying its integrity.

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed on your machine.
2. Clone the repository:
   ```bash
   git clone https://github.com/crncevicmarko/simpleBlockchain.git
   cd simpleBlockchain

## How to Run

To run the blockchain program, make sure you are positioned in the `simpleBlockchain` directory in terminal. Then, use the following command:

```bash
go run main.go block.go profeOfWork.go
