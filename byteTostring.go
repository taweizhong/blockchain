// 字节转字符串
package main

import "fmt"

func main() {
	var b []byte = []byte{'1', '2', '3', '4'}
	fmt.Printf("%T %v", string(b), string(b))
}
