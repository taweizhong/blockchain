// 字节转整型
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var n int32 = 20180516   
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, n)
	fmt.Printf("%T %v", buf.Bytes(), buf.Bytes())
}
