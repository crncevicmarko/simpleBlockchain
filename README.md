# Simple Blockchain Implementation in Go

This repository contains a basic implementation of a blockchain in Go. The blockchain consists of blocks that store data, a timestamp, and a reference to the previous block's hash. The main goal of this project is to demonstrate the foundational concepts of blockchain technology and how to create a simple blockchain structure in Go.

## Features

- **Block Structure**: Each block contains:
  - `TimeStamp`: The time the block was created.
  - `Data`: The information stored in the block.
  - `PrevBlockHash`: A hash of the previous block, ensuring the integrity of the chain.
  - `Hash`: The hash of the current block, calculated based on its contents.

- **Blockchain Structure**: 
  - A blockchain consists of multiple blocks linked together.
  - The first block is known as the Genesis block.

- **Hashing**: 
  - Uses the SHA-256 hashing algorithm to generate a hash for each block.
  - The hash is calculated from the block's timestamp, data, and the previous block's hash.

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed on your machine.
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/simple-blockchain.git
   cd simple-blockchain

