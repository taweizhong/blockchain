// 获取命令参数
package main

import (
	"flag"
	"fmt"
	"os"
)

type person struct {
	name string
	age  string
}

func useArgs() {
	args := os.Args
	p := person{
		name: args[1],
		age:  args[2],
	}
	fmt.Printf("%+v", p)
}

func useFlag() {
	var Name string
	var Age string
	flag.StringVar(&Name, "name", "姓名", "")
	flag.StringVar(&Age, "age", "年龄", "")
	flag.Parse()
	p := person{
		name: Name,
		age:  Age,
	}

	fmt.Printf("%+v", p)
}

func main() {
	// useArgs()
	useFlag()
}
