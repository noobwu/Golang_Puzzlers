package main

import "fmt"

type Pet interface {

	// @title    函数名称
	// @description   函数的详细描述
	// @auth      作者             时间（2019/6/18   10:57 ）
	// @param     输入参数名        参数类型         "解释"
	// @return    返回参数名        参数类型         "解释"

	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

// @title    函数名称
// @description   函数的详细描述
// @auth      作者             时间（2019/6/18   10:57 ）
// @param     输入参数名        参数类型         "解释"
// @return    返回参数名        参数类型         "解释"

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
	dog := Dog{"little pig"}

	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)

	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}
