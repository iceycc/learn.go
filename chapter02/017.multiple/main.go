package main

import (
	"fmt"
)

// init 函数
// 1. init 函数可以有多个
// 2. init 函数在 main 函数之前自动执行
func init() {
	fmt.Println("init")
}

func main() {
	// bmis := []float64{1.2, 3.222, 4.343443}
	avgBmi := calculateAvg(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(avgBmi)
	avgBmi = calculateAvgOnSlice([]float64{})
	fmt.Println(avgBmi)
	getScoresOfStudent("Tom")

	getNames(func(name string, age int) {
		fmt.Println("getNames 1", name, age)
	})
	getNameFns := getNamesFn()
	getNameFns("Tom", 18)

	var value = func() int {
		return 100
	}
	value()
}

// 函数可以当作参数传入
func getNames(fn func(name string, age int)) {
	fn("Tom", 18)
}

// 函数也可以作为返回值
func getNamesFn() func(name string, age int) {
	return func(name string, age int) {
		fmt.Println("getNames 2", name, age)
	}
}

func calculateAvg(bmis ...float64) (avgBmi float64) {
	fmt.Println("bmis-----", bmis)
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis))
	return
}

func calculateAvgOnSlice(bmis []float64) float64 {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}

func getScoresOfStudent(stdName string) (chinese int, match int, english int, physics int, nature int) {
	return 0, 0, 0, 0, 0
}
