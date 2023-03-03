// go中，每个可运行程序都有一个主package。他用来告诉go这个是个可运行的程序
package main

import (
	"fmt"
)

// main函数是程序的入口。程序运行时的起点。
func main() {
	fmt.Println("1、打开冰箱")
	fmt.Println("2、把大象放进去")
	fmt.Println("3、关上冰箱")
}
