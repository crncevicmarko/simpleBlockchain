package main

import (
	"log"
	"time"
)

type BlockChain struct {
	storage []*Block // like the database
}

type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0} // create new instance of the object Block with atributes
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	log.Print("Pow: ", pow)

	return block // return new objeck with set hash value
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlockHash := bc.storage[len(bc.storage)-1] // getting the last element in the list
	newBlock := NewBlock(data, prevBlockHash.Hash) // new block
	bc.storage = append(bc.storage, newBlock)      // append new block to database
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // first block in the list
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
