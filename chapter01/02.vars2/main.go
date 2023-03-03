package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "小强"
	{
		val, err := strconv.Atoi(name)
		fmt.Println(val, err)
	}
	age := 30
	// 打印变量的地址
	fmt.Printf("%p\n", &age)
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age)
	fmt.Println(name, age, time)
	// 同一个作用域中，不能重复定义变量
	{
		age := 3
		fmt.Printf("%p\n", &age)
	}
}
