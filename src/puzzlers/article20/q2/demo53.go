package main

import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
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
	flag.Parse()
	greeting, err := hello(name)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println(greeting, introduce())
}

// hello 用于生成问候内容。
func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

// introduce 用于生成介绍内容。
func introduce() string {
	return "Welcome to my Golang column."
}
