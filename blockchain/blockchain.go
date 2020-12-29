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

// Initializes a blockchain structure
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
