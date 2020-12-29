package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

// const (
// 	dbPath = "./tmp/blocks"
// )

// Struct that represents a blockchain structure
type BlockChain struct {
	LastHash []byte     // hash of the newest block in the blockchain
	Database *badger.DB // connection to the badger database
	// Blocks []*Block
}

// Member function used for adding a new block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, intErr := txn.Get([]byte("lh"))
		handleErr(intErr)
		lastHash, intErr = item.ValueCopy(nil)

		return intErr
	})
	handleErr(err)

	newBlock := CreateBlock(data, lastHash)
	err = chain.Database.Update(func(txn *badger.Txn) error {
		intErr := txn.Set(newBlock.Hash, newBlock.Serialize())
		handleErr(intErr)
		intErr = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return intErr
	})
	handleErr(err)
}

// Creates an iterator for a blockchain
func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

// Initializes a blockchain structure and saves it to the given path
func InitBlockChain(dbPath string) *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	handleErr(err)

	err = db.Update(func(txn *badger.Txn) error {
		// lh - last hash
		if _, intErr := txn.Get([]byte("lh")); intErr == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Created and proved genesis block")

			intErr = txn.Set(genesis.Hash, genesis.Serialize())
			intErr = txn.Set([]byte("lh"), genesis.Hash) // only block in database so set it as last hash
			lastHash = genesis.Hash

			return intErr
		} else {
			item, intErr := txn.Get([]byte("lh"))
			handleErr(intErr)
			// lastHash both on left and right side because if capacity of dst isn't sufficient,
			// a new slice would be allocated and returned (TODO: Check if this won't return nil)
			lastHash, intErr = item.ValueCopy(nil)
			return intErr
		}
	})
	handleErr(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

// Simple iterator for the blockchain structure
type BlockChainIterator struct {
	CurrentHash []byte     // hash of the current iteration
	Database    *badger.DB // connection to the badger database
}

// Returns a pointer to the next block
func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, intErr := txn.Get(iter.CurrentHash)
		handleErr(intErr)
		encodedBlock, intErr := item.ValueCopy(nil) // TODO: Check if this nil will work as intended
		block = Deserialize(encodedBlock)

		return intErr
	})
	handleErr(err)

	iter.CurrentHash = block.PrevHash
	return block
}
