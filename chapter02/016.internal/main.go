package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()
	// deffer 会在函数结束时执行
	// 会按照先进后出的顺序执行
	// 会在return、panic、之前执行

	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println(finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	arr2 := testPanic()

	arr3 := make([]string, 3, 4)
	copy(arr3, arr2)

	i := 3333
	// 通过指针获取值
	j := &i
	fmt.Println(reflect.TypeOf(j))

	//
	arr4 := []int{}
	fmt.Println(arr4)
	// make 会初始化数组
	// new
	arr5 := make([]string, 0, 4)
	fmt.Println(arr5)

	i2 := new(int)
	fmt.Println(i2)

}

func testPanic() []string {
	arr2 := make([]string, 1, 4)
	fmt.Println("len: ", len(arr2), "cap: ", cap(arr2))
	fmt.Println("default:", arr2[0])
	fmt.Println("default:", arr2[1])
	fmt.Println("default:", arr2[2])
	return arr2
}
