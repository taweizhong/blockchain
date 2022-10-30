// 加密算法
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func MD5() {
	str := "知链区块链"
	hash := md5.Sum([]byte(str))
	fmt.Printf("%x", hash)
}

func SHA1() {
	str := "知链区块链"
	hash := sha1.Sum([]byte(str))
	fmt.Printf("%x", hash)
}

func SHA2() {
	str := "知链区块链"
	hash := sha256.Sum256([]byte(str))
	fmt.Printf("%x", hash)
}

func main() {
	// MD5()
	// SHA1()
	SHA2()
}
