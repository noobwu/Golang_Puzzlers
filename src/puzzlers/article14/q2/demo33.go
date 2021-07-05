package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	var dog1 *Dog
	fmt.Println("The first dog is nil.")
	dog2 := dog1
	fmt.Println("The second dog is nil.")
	/*
		当我们把dog2的值赋给变量pet的时候，dog2的值会先被复制，不过由于在这里它的值是nil，所以就没必要复制了。

		然后，Go 语言会用我上面提到的那个专用数据结构iface的实例包装这个dog2的值的副本，这里是nil。

		虽然被包装的动态值是nil，但是pet的值却不会是nil，因为这个动态值只是pet值的一部分而已。

		顺便说一句，这时的pet的动态类型就存在了，是*Dog。我们可以通过fmt.Printf函数和占位符%T来验证这一点，另外reflect包的TypeOf函数也可以起到类似的作用。
	*/
	var pet Pet = dog2
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
	fmt.Printf("The type of pet is %T.\n", pet)
	fmt.Printf("The type of pet is %s.\n", reflect.TypeOf(pet).String())
	fmt.Printf("The type of second dog is %T.\n", dog2)
	fmt.Println()

	// 示例2。
	wrap := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}
	pet = wrap(dog2)
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
}
