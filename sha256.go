package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "123456"
	hash := sha256.Sum256([]byte(str))
	fmt.Printf("%T %x\n", hash[:], hash[:])

	fmt.Printf("%v", hex.EncodeToString(hash[:]))

}
