package main

import (
	"bytes"         // need to convert data into byte in order to be sent on the network, computer understands better the byte(8bits)language
	"crypto/sha256" //crypto library to hash the data
	"strconv"       // for conversion
	"time"          // use timestamp
)

// Generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                    // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.BlockData}, []byte{}) // concatenate all the block data
	hash := sha256.Sum256(headers)
	block.BlockHash = hash[:]
}

// Block Generation and return that block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)} // the block received
	block.SetHash()                                                           // the block is hashed
	return block                                                              // return the block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // Creating genesis block with some data
}
