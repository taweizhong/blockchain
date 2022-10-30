// 读文件
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		str, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}
