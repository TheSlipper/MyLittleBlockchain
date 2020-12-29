package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

// A simple error handling function
func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Casts a 64 bit integer to an array of bytes
func toHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	handleErr(err)

	return buff.Bytes()
}
