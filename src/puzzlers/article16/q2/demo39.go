package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {
	num := 10
	sign := make(chan struct{}, num) //struct{}:既不包含任何字段也不拥有任何方法的空结构体类型

	for i := 0; i < num; i++ {
		go func() {
			fmt.Printf("i:%d,currentThreadId:%d\n", i, windows.GetCurrentThreadId())
			/*
				注意，struct{}类型值的表示法只有一个，即：struct{}{}。并且，它占用的内存空间是0字节。确切地说，
				这个值在整个 Go 程序中永远都只会存在一份。虽然我们可以无数次地使用这个值字面量，但是用到的却都是同一个值。
			*/
			sign <- struct{}{}
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500) //休眠

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
}
