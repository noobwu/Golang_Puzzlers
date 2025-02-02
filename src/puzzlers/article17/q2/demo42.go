package main

import "fmt"

func main() {
	// 示例1。
	//value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	/*
		由于上述代码中的switch表达式的结果类型是int，而那些case表达式中子表达式的结果类型却是int8，
		它们的类型并不相同，所以这条switch语句是无法通过编译的。
	*/
	//switch 1 + 3 { // 这条语句无法编译通过。
	//case value1[0], value1[1]:
	//	fmt.Println("0 or 1")
	//case value1[2], value1[3]:
	//	fmt.Println("2 or 3")
	//case value1[4], value1[5], value1[6]:
	//	fmt.Println("4 or 5 or 6")
	//}

	// 示例2。
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6} //数组类型

	/*
		如果case表达式中子表达式的结果值是无类型的常量，那么它的类型会被自动地转换为switch表达式的结果类型，
		又由于上述那几个整数都可以被转换为int8类型的值，所以对这些表达式的结果值进行判等操作是没有问题的。
	*/
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}
