package blockchain

import "math/big"

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements

// Requirements:
// The first few bytes must contain 0s

// This implementation of blockchain has a static difficulty however in real-life appliances one would
// have an algorithm that would incrementally increase the difficulty over a large period of time.
// Main reason to do that is to account for the increase of miners on the network and account for
// increase in computation power of computers in general
const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join (
		[][]byte {
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{}
	)
	return data
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) // 256 - number of bytes inside one of our hashes

	pow := &ProofOfWork{b, target}
	return pow
}

// Casts a 64 bit integer to an array of bytes
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
