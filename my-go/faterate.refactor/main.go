package main

import (
	"fmt"
	ICalc "learn.go/my-go/faterate.refactor/ICalc"
)

// 计算体脂
// 计算体脂率
// 体脂率 = 体脂重量 / 体重

func main() {
	bmi2, err := ICalc.CalcBMI2(1.7, 60)
	if err != nil {
		return
	}
	fmt.Println(bmi2)
}
