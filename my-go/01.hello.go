package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//变量
	var hello string = "hello , golang"
	// 类型推断
	fmt.Println(reflect.TypeOf(hello))
	var world = "world"
	fmt.Println(hello, world)
	var age int = 1
	fmt.Println(age)
	var tall float32 = 1.1
	fmt.Println(tall)
	//tall = "22121"
	var xs = 1.22
	fmt.Println(xs)
	小数 := 1.1
	fmt.Println(小数)

	var l1, l2 uint = 1, 2 // 猜的和预期的不一样
	fmt.Println(l1, l2)

	// 猜测的类型和期望的不一致就会报错
	var co1, co2 float64 = 1, 2.2
	fmt.Println(co1 * co2)

	// 类型推断
	var ao int
	fmt.Println(ao)

	// 类型转换，丢失精度
	var j1, j2 = 1, 2.2
	fmt.Println(j1 + int(j2))

	// 类型转换, 小数转整数，存储方式不同，导致精度丢失严重
	var int6 uint = math.MaxUint64
	fmt.Println(int6) // 18446744073709551615
	var int7 int = int(int6)
	fmt.Println(int7) // -1 // 原码、反码、补码

}
