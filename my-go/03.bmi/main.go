package main

import "fmt"

// BMI= 体重(kg) / 身高(m)的平方

func getBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

// 计算体脂率
// 体脂率 = (1.2 * BMI) + (0.23 * 年龄) - (10.8 * 性别) - 5.4
// 性别：男=1，女=0
/**
 * @param bmi
 * @param age
 * @gender int 1:男 0:女
 */
func getBmiRate(bmi float64, age int, gender int) float64 {
	return (1.2*bmi + 0.23*float64(age) - 10.8*float64(gender) - 5.4) / 100
}
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
func main() {
	var weight float64 = 80
	var height float64 = 1.7
	var age int = 31
	var six string = "男"
	var myBmi float64 = getBMI(weight, height)
	var myBmiRate float64 = getBmiRate(myBmi, age, If(six == "男", 1, 0).(int))
	fmt.Println("请输入体重(kg)：", myBmi, myBmiRate)

}
