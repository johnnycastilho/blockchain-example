package main

// Add a new block to the blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // get the previous block
	newBlock := NewBlock(data, previousBlock.BlockHash)          // create the new block with the data and the previous block
	blockchain.Blocks = append(blockchain.Blocks, newBlock)      // add the block to the chain to create a chain of blocks
}

// Create the function that returns the whole blockchain and add the genesis to it first.
// the genesis block is the first ever mined block, so let's create a function that will return it since it does not exist yet
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
