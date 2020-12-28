package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Struct that represents a blockchain structure
type BlockChain struct {
	blocks []*Block
}

// Member function used for adding a new block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Struct that represents a single block in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Creates a hash for the invoking block and stores it inside the block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // this is fairly similar to real life hash calculating in blockchain but not secure enough
	b.Hash = hash[:]
}

// Creates a block that can be appended to the block with the prevHash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Creates the first block in the blockchain - the genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Initializes a blockchain structure
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	// add some more blocks
	chain.AddBlock("A->B;200") // A transfers 200 units to B
	chain.AddBlock("B->A;50")
	chain.AddBlock("B->C;330")

	for _, block := range chain.blocks {
		fmt.Printf("\n")
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data contained in this Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
