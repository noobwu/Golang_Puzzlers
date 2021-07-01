package main

import (
	"flag" //首先，Go 语言标准库中有一个代码包专门用于接收和解析命令参数。这个代码包的名字叫flag。
	"fmt"
	"os"
)

var name string

// 方式3。
//var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 方式2。
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

	/*
		自定义命令源码文件的参数使用说明
		注意，对flag.Usage的赋值必须在调用flag.Parse函数之前。
	*/
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	// 方式3。
	//cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
	/*
		函数flag.StringVar接受 4 个参数
			第 1 个参数是用于存储该命令参数值的地址，具体到这里就是在前面声明的变量name的地址了，由表达式&name表示。
			第 2 个参数是为了指定该命令参数的名称，这里是name。
			第 3 个参数是为了指定在未追加该命令参数时的默认值，这里是everyone。
			至于第 4 个函数参数，即是该命令参数的简短说明了，这在打印命令说明时会用到。
	*/
	flag.StringVar(&name, "name", "everyone", "The greeting object.")

}

func main() {
	// 方式1。
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	// 方式3。
	//cmdLine.Parse(os.Args[1:])
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
