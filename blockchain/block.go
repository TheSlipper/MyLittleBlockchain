package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Struct that represents a blockchain structure
type BlockChain struct {
	Blocks []*Block
}

// Member function used for adding a new block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
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
