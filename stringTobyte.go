// 字符串转字节
package main

import "fmt"

func main() {
	str := "123456"
	fmt.Printf("%T %v", []byte(str), []byte(str))
}
