package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

func main() {
	var a = 10
	fmt.Println(^a)
	a, b := 3, 4
	// 异或的交换律
	fmt.Println(a ^ b)
	fmt.Println(b ^ a)
	arr := []int{3, 4, 8, 3, 4, 8, 5, 9, 6, 5, 6}
	result := -1
	//
	for i, item := range arr {
		fmt.Println(i, result)
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(result)
	fmt.Printf("我的名字是：%s, 我的年龄是：%d\n", "小强", 30)
	sha256.New()
	var c1, c2 = 1, 1
	s1 := math.Pi * math.Pow(float64(c1), 2)
	s2 := math.Pi * math.Pow(float64(c2), 2)
	fmt.Printf("%.2f\n%.2f\n%.2f", s1, s2, s1+s2)
}
