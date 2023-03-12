package main

import (
	"fmt"
)

func main() {
	// 位运算的应用
	// 1. 交换两个变量的值
	// 2. 从一个数组中找出唯一一个不重复的元素
	// 3. 从一个数组中找出唯一一个重复的元素
	a, b := 100, 31
	fmt.Println(a ^ b)
	fmt.Println(b ^ a)

	arr := []int{4, 3, 4, 5, 6, 7, 3, 5, 6}
	result := -1
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(result)
}
