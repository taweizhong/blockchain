package main

import (
	"fmt"
	"os"
)

type OsCli struct {
}

func (o *OsCli) Init() {
	cads := os.Args
	switch cads[1] {
	case "-Creat":
		fmt.Println("创建区块链")
		bc := CreatBlockChain()
		if bc.Save(File) {
			fmt.Println("保存成功")
		}

	case "-add":
		fmt.Println("添加区块")
		var b BlockChain
		if b.Load(File) {
			fmt.Println("读取成功")
		}
		fmt.Println("------------------")
		b.AddBlock(cads[2])
		if b.Save(File) {
			fmt.Println("保存成功")
		}
	case "-P":
		fmt.Println("打印区块链")
		var b BlockChain
		b.Load(File)
		b.PrintBlockChain()
	}
}
