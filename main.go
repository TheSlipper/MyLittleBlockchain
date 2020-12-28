package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TheSlipper/MyLittleBlockchain/blockchain"
)

func main() {
	start := time.Now()

	chain := blockchain.InitBlockChain()

	// add some more blocks
	chain.AddBlock("A->B;200;CA:BRITISHCOLOMBIA") // A transfers 200 units to B
	chain.AddBlock("B->A;50;US:ALASKA")
	chain.AddBlock("B->C;330;EU:POLAND")

	for _, block := range chain.Blocks {
		fmt.Printf("\n")
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data contained in this Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\nNonce:%d\n", strconv.FormatBool(pow.Validate()), block.Nonce)
		fmt.Println()
	}

	// elapsed := time.Now().Sub(start)
	fmt.Printf("Time elapsed while calculating hashes and validating them: %v\n", time.Since(start))
}
