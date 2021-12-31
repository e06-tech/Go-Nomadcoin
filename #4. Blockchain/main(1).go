package main

import "fmt"
import "crypto/sha256"

type block struct {
	data string
	hash string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
} // %x: convert hexadecimal numbers to strings and numerals