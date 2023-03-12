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
func scan() (string, float64, float64) {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重(kg)：")
	fmt.Scanln(&weight)

	var height float64
	fmt.Print("身高(m)：")
	fmt.Scanln(&height)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var six string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&six)

	var myBmi float64 = getBMI(weight, height)
	var myBmiRate float64 = getBmiRate(myBmi, age, If(six == "男", 1, 0).(int))
	fmt.Println(myBmi, myBmiRate)
	if myBmiRate < 0.1 {
		fmt.Println("偏瘦")
	} else if myBmiRate < 0.2 {
		fmt.Println("正常")
	} else if myBmiRate < 0.3 {
		fmt.Println("超重")
	} else {
		fmt.Println("肥胖")
	}
	var flag string
	fmt.Println("是否继续(y/n)：")
	fmt.Scanln(&flag)
	return flag, myBmi, myBmiRate
	//if flag == "y" {
	//	scan()
	//} else {
	//	return
	//}
}
func main() {
	var bmis [3]float64
	var bmiRates [3]float64
	var i = 0
	for {
		var flag, myBmi, myBmiRate = scan()
		bmis[i] = myBmi
		bmiRates[i] = myBmiRate
		if flag == "n" {
			break
		}
		i++
		if i == 3 {
			break
		}
	}
	var sumBmi float64 = 0
	for i := 0; i < len(bmis); i++ {
		sumBmi += bmis[i]
	}
	println(int(sumBmi / float64(len(bmis))))
}
