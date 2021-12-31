package main

import "fmt"
import "crypto/sha256"

type block struct {
	data string
	hash string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain

func GetBlockchain()*blockchain {
	if b == nill {
		b = &blcokchain{}
	}
	retrun b
}