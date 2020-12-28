package blockchain

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

// Block represents a single blockchain block
type Block struct {
	Hash     []byte // Hash of the information residing in this block
	Data     []byte // Data stored in this block structure
	PrevHash []byte // Hash of the previous block
	Nonce    int
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

// Initializes a blockchain structure
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
