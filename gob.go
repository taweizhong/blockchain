// gob文件读写
package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

func w() {
	str := "hello, golang\n 你好， 世界"
	file, _ := os.OpenFile("./test.gob", os.O_CREATE|os.O_RDWR, 0666)
	buf := bufio.NewWriter(file)
	defer file.Close()
	en := gob.NewEncoder(buf)
	en.Encode(str)
	buf.Flush()
}

func r() {
	var str string
	f, _ := ioutil.ReadFile("./test.gob")
	re := bytes.NewBuffer(f)
	de := gob.NewDecoder(re)
	de.Decode(&str)
	fmt.Println(str)
}

func main() {
	w()
	r()
}
