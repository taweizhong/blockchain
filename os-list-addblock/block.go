package main

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
	"time"
)

const Version = 1.0
const Difficulty = 1

type Block struct {
	Hash       []byte
	PHash      []byte
	ChainH     int
	Version    int
	MerKleRoot string
	TimeStamp  int64
	Nonce      int
	Difficulty int
	Data       string
}

func NewBlock(data string, phash []byte, high int) *Block {
	block := Block{
		PHash:      phash,
		ChainH:     high,
		Version:    Version,
		MerKleRoot: "",
		TimeStamp:  time.Now().Unix(),
		Nonce:      rand.Intn(500),
		Difficulty: Difficulty,
		Data:       data,
	}
	block.setHash()
	return &block
}

func (block *Block) setHash() {
	data := block.Data + strconv.FormatInt(block.TimeStamp, 10) + strconv.Itoa(block.Nonce)
	hash := sha256.Sum256([]byte(data))
	block.Hash = hash[:]
}
