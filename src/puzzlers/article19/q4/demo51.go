package main

import "fmt"

func main() {
	/*
		在defer语句每次执行的时候，Go 语言会把它携带的defer函数及其参数值另行存储到一个链表中。

		这个链表与该defer语句所属的函数是对应的，并且，它是先进后出（FILO）的，相当于一个栈。

		在需要执行某个函数中的defer函数调用的时候，Go 语言会先拿到对应的链表，然后从该链表中一个一个地取出defer函数及其参数值，并逐个执行调用。

		这正是我说“defer函数调用与其所属的defer语句的执行顺序完全相反”的原因了。
	*/
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}
