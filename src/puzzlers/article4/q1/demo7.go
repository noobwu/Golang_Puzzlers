package main

import (
	"flag" //首先，Go 语言标准库中有一个代码包专门用于接收和解析命令参数。这个代码包的名字叫flag。

	"fmt"
)

func main() {
	var name string // [1]

	/*
		函数flag.StringVar接受 4 个参数
			第 1 个参数是用于存储该命令参数值的地址，具体到这里就是在前面声明的变量name的地址了，由表达式&name表示。
			第 2 个参数是为了指定该命令参数的名称，这里是name。
			第 3 个参数是为了指定在未追加该命令参数时的默认值，这里是everyone。
			至于第 4 个函数参数，即是该命令参数的简短说明了，这在打印命令说明时会用到。

	*/

	/*
		注意，flag.String函数返回的结果值的类型是*string而不是string。类型*string代表的是字符串的指针类型，而不是字符串类型。
		因此，这里的变量name代表的是一个指向字符串值的指针。
	*/
	flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]

	// 方式1。
	//var name = flag.String("name", "everyone", "The greeting object.")

	// 方式2。
	//name := flag.String("name", "everyone", "The greeting object.")

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)

	// 适用于方式1和方式2。
	//fmt.Printf("Hello, %v!\n", *name)
}
