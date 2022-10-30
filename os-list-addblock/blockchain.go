package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
)

type BlockChain struct {
	blockChain []*Block
}

func CreatBlockChain() *BlockChain {
	firstBlock := NewBlock("创世区块", nil, 0)
	myBlockChain := BlockChain{
		blockChain: []*Block{
			firstBlock,
		},
	}
	return &myBlockChain
}

func (this *BlockChain) AddBlock(data string) {
	pHash := this.blockChain[len(this.blockChain)-1].Hash
	high := len(this.blockChain)
	block := NewBlock(data, pHash, high)
	this.blockChain = append(this.blockChain, block)
}

func (this *BlockChain) PrintBlockChain() {
	for i, v := range this.blockChain {
		fmt.Printf("%d\n", i)
		fmt.Printf("区块高度：%d\n", v.ChainH)
		fmt.Printf("区块时间戳：%d\n", v.TimeStamp)
		fmt.Printf("区块随机数：%d\n", v.Nonce)
		fmt.Printf("区块hash:%x\n", v.Hash)
		fmt.Printf("区块内部数据：%s\n", v.Data)
		fmt.Println("==================")
	}
}

func (this *BlockChain) Save(file string) bool {
	buf := bytes.Buffer{}
	gob.Register(elliptic.P256())
	en := gob.NewEncoder(&buf)
	err := en.Encode(this.blockChain)
	if err != nil {
		return false
	}
	err = ioutil.WriteFile(file, buf.Bytes(), 0600)
	if err != nil {
		return false
	}
	return true
}

func (this *BlockChain) Load(file string) bool {
	read, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	gob.Register(elliptic.P256())
	de := gob.NewDecoder(bytes.NewBuffer(read))
	var block []*Block
	err = de.Decode(&block)
	if err != nil {
		return false
	}
	this.blockChain = block
	return true
}
