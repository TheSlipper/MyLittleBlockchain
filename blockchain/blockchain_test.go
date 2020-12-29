package blockchain

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

const (
	dbPath = "./tmp/testing_bc"
)

// Cleans up the blockchain files from the disk
func cleanUpDb(t *testing.T) {
	var dbfiles []string

	err := filepath.Walk(dbPath, func(path string, info os.FileInfo, err error) error {
		dbfiles = append(dbfiles, path)
		return nil
	})

	if err != nil {
		t.Log("Could not retrieve the database file names")
		t.Fail()
	}

	for _, file := range dbfiles {
		err = os.Remove(file)
	}
}

// Populates the blockchain with entries
func populateBlockchain(chain *BlockChain) {
	chain.AddBlock("A->B;200;CA:BRITISHCOLOMBIA")
	chain.AddBlock("B->A;50;US:ALASKA")
	chain.AddBlock("B->C;330;EU:POLAND")
}

// Tests blockchain creation, blockchain hash continuity, proof of work and validation
func TestBlockchain(t *testing.T) {
	var prevHash []byte
	chain := InitBlockChain(dbPath)
	defer cleanUpDb(t)
	populateBlockchain(chain)

	iter := chain.Iterator()

	// Extract last block from iterator
	block := iter.Next()
	prevHash = block.PrevHash
	pow := NewProof(block)
	if !pow.Validate() {
		t.Log("Proof of work for genesis block failed")
		t.Fail()
	}

	for block.PrevHash != nil {
		block = iter.Next()
		if !bytes.Equal(block.Hash, prevHash) {
			t.Logf("prevHash: '%x', block.Hash: %x\n", prevHash, block.Hash)
			t.Log("Block.PrevHash is not equal to Block.Hash in the previous block")
			t.Fail()
		}

		pow = NewProof(block)
		if !pow.Validate() {
			t.Log("Proof of work for genesis block failed")
			t.Fail()
		}

		prevHash = block.PrevHash
	}
}
