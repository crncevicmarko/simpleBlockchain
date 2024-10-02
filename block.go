package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

//	we need to set hash value for the bloch entity
//
// we will do it if we set timestamp, data and prevBlockHash atributtes in header and we hash it
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))                       // convert timestamp to byte slice(like string of digits)
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // join all atributes in headers variable
	hash := sha256.Sum256(headers)                                                // hash header with sum256 func

	b.Hash = hash[:] // set hash value to Hash atribute
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}} // create new instance of the object Block with atributes
	block.SetHash()                                                           // set hash value for object Block
	return block                                                              // return new objeck with set hash value
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
