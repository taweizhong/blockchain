package main

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
	"time"
)

// 哈希
// 前置哈希
// 区块高度
// 版本号
// 默克尔树
// 时间戳
// 难度值
// 随机数
// 数据
type Block struct {
	Hash          []byte
	PrevBlockHash []byte
	ChainHigh     int
	Version       int
	MerkleRoot    string
	TimeStamp     int64
	Difficulty    int
	Nonce         int
	Data          string
}

func NewBlock(data string, prevblockchain []byte, chainhigh int) *Block {
	block := Block{
		ChainHigh:     chainhigh,
		PrevBlockHash: prevblockchain,
		Nonce:         rand.Intn(500),
		Version:       Version,
		MerkleRoot:    "",
		TimeStamp:     time.Now().Unix(),
		Difficulty:    0,
		Data:          data,
	}
	block.setHash()
	return &block
}

func (T *Block) setHash() {
	data := strconv.FormatInt(T.TimeStamp, 10) + string(rune(T.Nonce)) + T.Data
	hash := sha256.Sum256([]byte(data))
	T.Hash = hash[:]
}
