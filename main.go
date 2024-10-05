package main

import (
	"log"
	"strconv"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.storage {
		log.Printf("Prev hash: %x\n", block.PrevBlockHash)
		log.Printf("Data: %s\n", block.Data)
		log.Printf("Hash: %x\n", block.Hash)
		log.Printf("Nonce: %x\n", block.Nonce)
		pow := NewProofOfWork(block)
		log.Printf("PoW: %s\n", strconv.FormatBool(pow.Validation()))
		log.Println()
	}
}
