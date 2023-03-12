package main

import (
	"fmt"
	"reflect"
)

func main() {
	//leftMap, rightMap := map[string]int{}, map[string]int{}
	//
	//leftMap["语文"] = 80
	//rightMap["数学"] = 43
	//
	//for k, v := range rightMap {
	//	leftMap[k] = v
	//}
	//fmt.Println(leftMap)

	leftMap, rightMap := map[string]int{}, map[string]int{}

	leftMap["语文"] = 80
	rightMap["数学"] = 43
	for key, val := range leftMap {
		rightMap[key] = val
	}
	//fmt.Printf("%v", rightMap)

	// 函数传参数
	//var a = "a"
	//test(a)
	//println(a)

	var b = []string{"a", "b"}
	fmt.Println(b)
	test(b)
	fmt.Println(b)

}

func test(arg []string) {
	fmt.Println(reflect.TypeOf(arg))
	if reflect.TypeOf(arg) == reflect.TypeOf([]string{}) {
		arg[0] = "c"
	}
}
