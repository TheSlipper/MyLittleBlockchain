package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

/* ================================ PoW ================================
Proof of Work (PoW) algorithm steps:
1. Take the data from the block
2. Create a counter (nonce) which starts at 0 and increments upwards
theoretically infinitely
3. Create a hash of the data plus the counter
4. Check if the newly created hash meets a set of given requirements. If the hash meets the
requirements then the block is signed. If not then increment nonce and go back to step 3.

This PoW implementation has only one requirement: the first few bytes must contain 0s.
*/

// Defines the difficulty of signing a block. This implementation of blockchain has a static
// difficulty however in real-life appliances one would have an algorithm that would
// incrementally increase the difficulty over a large period of time. Main reason to do that
// is to account for the increase of miners on the network and account for increase in
// computation power of computers in general
const Difficulty = 16

// ProofOfWork contains a proof of work for a given block
type ProofOfWork struct {
	Block  *Block   // block that is meant to be signed
	Target *big.Int // represents the PoW requirement derived from the Difficulty
	// Perhaps the above is even the hash itself but in *big.Int format
	// TODO: Find out if that's true
}

// TODO: Consider to make the below private

// Initializes and returns the data in []byte format for further processing
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			toHex(int64(nonce)),
			toHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

// Runs the proof of work algorithm. Once it generates the desired hash it returns the nonce in
// the form of int and the hash in the form of []byte.
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		// This compares to the target. If intHash is smaller than pow.Target then it will
		// return -1. The goal of this is to check whether the requirement about the predefined
		// amount of zeros at the beginning of the hash is fulfilled
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()
	return nonce, hash[:]
}

// Validates whether this proof of work is valid by hashing it with the nonce present in the
// PoW struct's instance.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// Instantiates a new proof of work of the given block. It is important to note that this PoW
// is not yet ran therefore it is not signed.
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// This left shifts target bitwise by 244 (256-12)
	target.Lsh(target, uint(256-Difficulty)) // 256 - number of bytes inside one of our hashes

	pow := &ProofOfWork{b, target}
	return pow
}
