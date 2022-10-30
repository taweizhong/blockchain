// 写文件
package main

import (
	"bufio"
	"os"
)

func main() {
	str := "hello, golang"
	file, _ := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	buf := bufio.NewWriter(file)
	buf.WriteString(str)
	buf.Flush()
	defer file.Close()
}
