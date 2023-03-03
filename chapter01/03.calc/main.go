package main

import (
	"fmt"
)

func main() {
	// 数据类型选择不正确，计算
	var a, b int8 = 30, 11
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b) // 溢出。 玲玲，铃懿， 。。。。 忆初
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)

	fmt.Println(true && false == false)
	fmt.Println("a>b=", a > b)
	fmt.Println("a<b=", a < b)
}
