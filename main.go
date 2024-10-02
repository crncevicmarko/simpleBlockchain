package main

import (
	"fmt"
	"log"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Novi Blok 1")
	bc.AddBlock("Novi Blok 2")

	for _, block := range bc.storage {
		log.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		log.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
