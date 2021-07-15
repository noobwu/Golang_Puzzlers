package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		/*
			每条go语句一般都会携带一个函数调用，这个被调用的函数常常被称为go函数。
			而主 goroutine 的go函数就是那个作为程序入口的main函数。

			一定要注意，go函数真正被执行的时间，总会与其所属的go语句被执行的时间不同。
			当程序执行到一条go语句的时候，Go 语言的运行时系统，会先试图从某个存放空闲的 G 的队列中获取一个 G（也就是 goroutine），
			它只有在找不到空闲 G 的情况下才会去创建一个新的 G。

			请记住，只要go语句本身执行完毕，Go 程序完全不会等待go函数的执行，它会立刻去执行后边的语句。这就是所谓的异步并发地执行。
		*/
		go func() {
			fmt.Println(i) //不会有任何内容被打印出来
		}()
	}
}
