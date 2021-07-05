package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"} //数组类型

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"} //字典类型

	// 方式1。
	/*
		正式说明一下，类型断言表达式的语法形式是x.(T)。其中的x代表要被判断类型的值。这个值当下的类型必须是接口类型的，不过具体是哪个接口类型其实是无所谓的。

		这里有一条赋值语句。在赋值符号的右边，是一个类型断言表达式。

		它包括了用来把container变量的值转换为空接口值的interface{}(container)。

		以及一个用于判断前者的类型是否为切片类型 []string 的 .([]string)。

		这个表达式的结果可以被赋给两个变量，在这里由_和ok1代表。变量ok1是布尔（bool）类型的，它将代表类型判断的结果，true或false。

		如果是true，那么被判断的值将会被自动转换为[]string类型的值，并赋给变量value，否则value将被赋予nil（即“空”）。
	*/
	_, ok1 := interface{}(container).([]string) //类型断言表达式
	_, ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		container[1], container)

	// 方式2。
	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		elem, container)
}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return
	}
	return
}
