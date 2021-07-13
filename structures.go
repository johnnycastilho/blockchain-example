package main

type Block struct {
	Timestamp         int64  // the time when the block was created
	PreviousBlockHash []byte // the hash of previous block
	BlockHash         []byte // the hash of current block
	BlockData         []byte // transaction information [body of current block]
}

// Prepare blockchain data structure
type Blockchain struct {
	Blocks []*Block // Series of blocks
}
