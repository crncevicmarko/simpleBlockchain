package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 24 // how much 0 needs to be infront the hased value of block

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// we need func for the joining new block data
// we need to return the array of bytes(joined values from block data in form of bytes)
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// func for mining the correct hash value for the bloc and checking if the int representation of the hash value is the smaller of the target
// wee will checkout that with Cmp func that compares two int values one is the represantation of hash value of the block and other is the target number
// if the hash is less than target number Cmp func will return -1
// if is same it will return 0 and
// if its greater it will return 1

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte // sha256.Sum256(data) returns the [32]byte array
	nonce := 0

	log.Println("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		//we need for every value of nonce to check if the int representation of hash value of the block with that nonce gives the lesser value of target number

		//first we need to prepare block data
		data := pow.prepareData(nonce)
		//hash it
		hash = sha256.Sum256(data)
		//extract the int representation of the block hash
		hashInt.SetBytes(hash[:]) // hash[:] returns the slice of all elements of the hash

		//compare int values of hash and target
		//if hash is less than break
		//else increment the nonce
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

// we need func that will validate prof of walue of block(is this "hashInt.Cmp(pow.targer) == -1" returning true)
func (pow *ProofOfWork) Validation() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
