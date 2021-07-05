package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

// 方案1(函数作为传入参数)
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 方案2。
type calculateFunc func(x int, y int) (int, error)

//把函数作为结果返回
func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	// 方案1。
	x, y := 12, 23

	//匿名函数
	op := func(x, y int) int {
		return x + y
	}

	//把函数作为传入参数
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

	//把函数作为传入参数（函数参数为空）
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

	// 方案2(把函数作为结果返回)
	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
}
