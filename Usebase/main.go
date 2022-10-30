package main

import (
	"fmt"

	"github.com/mr-tron/base58"
)

func main() {
	str := "机密内容"

	en := base58.Encode([]byte(str))
	fmt.Printf("%v\n", en)

	de, _ := base58.Decode(en)
	fmt.Printf("%v", string(de))
}
