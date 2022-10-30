// 字符串转整型
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123456"
	n, _ := strconv.Atoi(str)
	fmt.Printf("%#v", n)
}
