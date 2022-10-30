package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/iocn-io/ripemd160"
)

func main() {
	str := "加密内容"

	hash := sha256.Sum256([]byte(str))

	h := ripemd160.New()
	h.Write(hash[:])
	hs := h.Sum(nil)

	fmt.Printf("%x", hs)

	

}
