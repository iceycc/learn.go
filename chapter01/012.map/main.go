package main

import (
	"fmt"
)

// map的注意事项
// 1. map的key必须是可以比较的，比如int、string、bool、struct、array、slice、map
func main() {
	names := []string{"小强", "周毅", "旭东"}
	fr := []float64{28, 18, 18}

	names = append(names, "张松")
	fr = append(fr, 19)

	for i, name := range names {
		if name == "周毅" {
			fmt.Printf("%s 的体脂率是 %f\n", name, fr[i])
		}
	}

	fmt.Println("定义Map")

	// nil 是空指针，不能直接使用
	var m1 map[string]int = nil
	// m1["a"] = 1 // panic on nil map
	delete(m1, "a")
	fmt.Println("m1 没有实例化，直接取数：", m1["a"])

	m2 := map[string][]float64{}
	m3 := map[string]int{"王强": 60, "李静": 83, "苗文": 99}
	fmt.Println(m1, m2, m3)

	fmt.Println("王强的分数：", m3["王强"])
	fmt.Println("王强的分数：", m3["李静"])
	fmt.Println("小强的分数：", m3["小强"])

	xqScore, ok := m3["小强"]
	fmt.Println(xqScore, ">>>>>", ok)

	m3["小强"] = 77
	fmt.Println("小强的分数：", m3["小强"])
	// map返回值有两个，第二个是bool类型，表示是否存在
	xqScore, ok = m3["小强"]
	fmt.Println(xqScore, ">>>>>", ok)

	for name, score := range m3 {
		fmt.Println(name, "=", score)
	}

	m4 := map[string]interface{}{
		"wby1": "小强",
		"age":  18,
	}
	fmt.Println(m4)

}

// 给10个同学的分数mock数据
