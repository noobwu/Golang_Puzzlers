package main

import "fmt"

func main() {
	/*
		一个通道相当于一个先进先出（FIFO）的队列。也就是说，通道中的各个元素值都是严格地按照发送的顺序排列的，
		先被发送通道的元素值一定会先被接收。元素值的发送和接收都需要用到操作符<-。我们也可以叫它接送操作符。
		一个左尖括号紧接着一个减号形象地代表了元素值的传输方向。
	*/
	ch1 := make(chan int, 3) //声明并初始化了一个元素类型为int、容量为3的通道ch1
	ch1 <- 2                 //向通道发送元素值
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1 //从通道接收元素值
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}
