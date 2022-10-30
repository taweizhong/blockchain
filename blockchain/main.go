package main

import "fmt"

func main() {
	B := CreatBlockChain() // 创建区块链，此时创世区块也一起创建好了
	fmt.Printf("%v\n", B)

	B.AddBlock("交易一:A 转账 B 10")
	B.AddBlock("交易一:A 转账 B 1010")
	B.AddBlock("交易一:A 转账 B 110")

	B.PrintBlock()

}
