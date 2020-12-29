package blockchain

import (
	"bytes"
	"encoding/gob"
)

// Block represents a single blockchain block
type Block struct {
	Hash     []byte // Hash of the information residing in this block
	Data     []byte // Data stored in this block structure
	PrevHash []byte // Hash of the previous block
	Nonce    int
}

// Returns a []byte representation of the block
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	handleErr(err)

	return res.Bytes()
}

// Creates a block from a []byte representation
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	handleErr(err)

	return &block
}

// Creates a block that can be appended to the block with the prevHash
func CreateBlock(data string, prevHash []byte) *Block {
	// Initialize the block with the data
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	// Instantiate a proof of work and then run it
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Creates the first block in the blockchain - the genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
