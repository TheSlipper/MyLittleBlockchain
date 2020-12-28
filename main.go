package main

import (
	"fmt"
	"github.com/TheSlipper/MyLittleBlockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	// add some more blocks
	chain.AddBlock("A->B;200") // A transfers 200 units to B
	chain.AddBlock("B->A;50")
	chain.AddBlock("B->C;330")

	for _, block := range chain.Blocks {
		fmt.Printf("\n")
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data contained in this Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
