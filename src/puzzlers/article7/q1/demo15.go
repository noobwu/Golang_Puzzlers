package main

import "fmt"

func main() {
	// 示例1。
	s1 := make([]int, 5) //初始化一个长度为5，容量为5的切片类型(可变数组)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)

	s2 := make([]int, 5, 8) //初始化一个长度为5，容量为8的切片类型(可变数组)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Println()

	// 示例2。
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	/*
		通用的规则是：一个切片的容量可以被看作是透过这个窗口最多可以看到的底层数组中元素的个数。

		由于s4是通过在s3上施加切片操作得来的，所以s3的底层数组就是s4的底层数组。

		又因为，在底层数组不变的情况下，切片代表的窗口可以向右扩展，直至其底层数组的末尾。

		所以，s4的容量就是其底层数组的长度8, 减去上述切片表达式中的那个起始索引3，即5。

		注意，切片代表的窗口是无法向左扩展的。也就是说，我们永远无法透过s4看到s3中最左边的那 3 个元素。
	*/

	s4 := s3[3:6]                                   //切片表达式s3[3:6]初始化了切片s4,[3:6]要表达的就是透过新窗口能看到的s3中元素的索引范围是从3到5（注意，不包括6）
	fmt.Printf("The length of s4: %d\n", len(s4))   //长度为3 索引[3,4,5]
	fmt.Printf("The capacity of s4: %d\n", cap(s4)) //容量为5
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println()

	// 示例3。
	s5 := s4[:cap(s4)]
	fmt.Printf("The length of s5: %d\n", len(s5))
	fmt.Printf("The capacity of s5: %d\n", cap(s5))
	fmt.Printf("The value of s5: %d\n", s5)
}
