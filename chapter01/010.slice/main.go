package main

import "fmt"

func main() {
	var a string = "您好"
	fmt.Println(a)

	// byte 和 rune
	// byte 是 uint8 的别名
	// rune 是 int32 的别名
	fmt.Println("byte a: ", []byte(a), len([]byte(a)))
	fmt.Println("rune a: ", []rune(a), len([]rune(a)))

	aBytes := []rune(a) // []int 不能转, []byte 可以和string互通
	fmt.Println(aBytes)
	fmt.Println("修改切片内的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)
}
