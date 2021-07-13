package main

// import the fmt package to

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("First Transaction")
	bc.AddBlock("Second Transaction")
	bc.AddBlock("Thirtd Transaction")

	for i, block := range bc.Blocks {
		fmt.Printf("Block ID: %d \n", i)
		fmt.Printf("Timestamp: %d \n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the block: %x\n\n", block.BlockHash)
		fmt.Printf("Hash of the previous block %x\n\n", block.PreviousBlockHash)
		fmt.Printf("All the transactions: %s\n\n", block.BlockData)
	}
}
