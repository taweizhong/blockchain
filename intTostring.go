// 整型转字符串
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int = 123456
	fmt.Printf("%#v", strconv.Itoa(n))
}
