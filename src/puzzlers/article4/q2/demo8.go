package main

import (
	"flag" //首先，Go 语言标准库中有一个代码包专门用于接收和解析命令参数。这个代码包的名字叫flag。
	"fmt"
)

func main() {
	var name = getTheFlag() //类型推断
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}

//上面函数的实现也可以是这样的。
//func getTheFlag() *int {
//	return flag.Int("num", 1, "The number of greeting object.")
//}
