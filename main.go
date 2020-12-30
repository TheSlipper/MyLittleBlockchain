package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheSlipper/MyLittleBlockchain/blockchain"
)

// Contains all arguments passed from the commandline
type runtimeArgs struct {
	action       string // defines what the program should do
	blockContent string // defines the content of a block that is to be added
	verbose      bool   // defines the amount of output to the console
	dbPath       string // path to the database with the blockchain data
}

// Parses the commandline arguments and populates a structure with them
func parseArgs() (args runtimeArgs) {
	flag.BoolVar(&args.verbose, "verbose", false, "when set to true more output will be produced")
	flag.StringVar(&args.dbPath, "input", "./tmp/db", "path to the database file")
	flag.Parse()

	// If there are no arguments left after parsing then it means that no actions where specified for
	// this command execution
	if flag.NArg() == 0 {
		flag.Usage()
		fmt.Print("\nACTION ARGUMENTS:\n  add string\n\tadds a block with the given message to")
		fmt.Println("the blockchain\n  print\n\tprints out all of the blockchain elements")
		os.Exit(1)
	} else { // If there are arguments however then parse them
		flagArgs := flag.Args()
		args.action = flagArgs[0]

		switch args.action {
		case "add":
			if len(flagArgs) < 2 {
				fmt.Printf("Insufficient amount of positional arguments! %s should have a message after it!\n",
					args.action)
				os.Exit(1)
			}
			args.blockContent = flagArgs[1]
		case "print":
		default:
			fmt.Printf("Invalid positional argument %s. Run with -help to see the command's usage\n",
				args.action)
			os.Exit(1)
		}
	}

	return
}

// Adds the block to the blockchain
func addBlock(args runtimeArgs) {
	chain := blockchain.InitBlockChain(args.dbPath)
	chain.AddBlock(args.blockContent)
}

// Prints out all of the blocks in the blockchain
func printBlocks(args runtimeArgs) {
	chain := blockchain.InitBlockChain(args.dbPath)
	iter := chain.Iterator()

	block := iter.Next()
	for {
		fmt.Printf("\n")
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data contained in this Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		if block.PrevHash == nil {
			break
		} else {
			block = iter.Next()
		}
	}
}

func main() {
	// Parse all of the arguments
	args := parseArgs()

	// Execute specified command
	switch args.action {
	case "add":
		addBlock(args)
	case "print":
		printBlocks(args)
	}
}
