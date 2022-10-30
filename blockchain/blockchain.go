package main

import "fmt"

type BlockChain struct {
	blocks []*Block
}

const Version = 1.0
const FirstData = "创世区块"

func CreatBlockChain() *BlockChain {
	firstblock := NewBlock(FirstData, nil, 0)
	BC := BlockChain{
		blocks: []*Block{
			firstblock,
		},
	}
	return &BC
}

func (T *BlockChain) AddBlock(data string) {
	prevblock := T.blocks[len(T.blocks)-1]
	prevblockhash := prevblock.Hash
	chainhigh := len(T.blocks)
	block := NewBlock(data, prevblockhash, chainhigh)
	T.blocks = append(T.blocks, block)
}

func (T *BlockChain) PrintBlock() {
	for i, block := range T.blocks {
		fmt.Printf("%d\n", i)
		fmt.Printf("区块高度：%d\n", block.ChainHigh)
		fmt.Printf("区块时间戳：%d\n", block.TimeStamp)
		fmt.Printf("区块随机数：%d\n", block.Nonce)
		fmt.Printf("区块hash:%x\n", block.Hash)
		fmt.Printf("区块内部数据：%s\n", block.Data)
		fmt.Println("==================")
	}
}
