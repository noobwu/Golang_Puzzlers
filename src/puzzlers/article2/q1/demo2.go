package main

import (
	"flag" //首先，Go 语言标准库中有一个代码包专门用于接收和解析命令参数。这个代码包的名字叫flag。
	"fmt"
)

var name string

func init() {

	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
