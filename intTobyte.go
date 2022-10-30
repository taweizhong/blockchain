// 整型转字节
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var b []byte = []byte{1, 51, 238, 36}
	var n int32
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &n)
	fmt.Printf("%T %v", n, n)
}
