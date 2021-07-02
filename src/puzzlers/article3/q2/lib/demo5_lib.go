package lib //作用域(相当于命名空间)

import "fmt" //引用系统文件

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
