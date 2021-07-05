package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 示例1。
	// 只能发不能收的通道。
	var uselessChan = make(chan<- int, 1)

	// 只能收不能发的通道。
	var anotherUselessChan = make(<-chan int, 1)

	// 这里打印的是可以分别代表两个通道的指针的16进制表示。
	fmt.Printf("The useless channels: %v, %v\n",
		uselessChan, anotherUselessChan)

	// 示例2。
	intChan1 := make(chan int, 3)
	//向单通道发送元素随机值
	/*
		顺便说一下，我们在调用SendInt函数的时候，只需要把一个元素类型匹配的双向通道传给它就行了，
		没必要用发送通道，因为 Go 语言在这种情况下会自动地把双向通道转换为函数所需的单向通道。
	*/
	SendInt(intChan1)

	// 示例4。
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	// 示例5。
	_ = GetIntChan(getIntChan)
}

//只能发不能收的单通道
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000) //向通道发送元素随机值
}

// 示例3。
type Notifier interface {
	SendInt(ch chan<- int)
}

//只能收不能发的单通道
/*
函数getIntChan会返回一个<-chan int类型的通道，这就意味着得到该通道的程序，只能从通道中接收元素值。
这实际上就是对函数调用方的一种约束了。
*/
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// 示例5。
type GetIntChan func() <-chan int
