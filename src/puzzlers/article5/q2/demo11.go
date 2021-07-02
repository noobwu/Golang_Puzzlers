package main

import "fmt"

var container = []string{"zero", "one", "two"} //string[] 类型

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"} //map类型
	fmt.Printf("The element is %q.\n", container[1])
}
