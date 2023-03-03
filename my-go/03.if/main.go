package main

import "fmt"

func main() {
	var fruit string = "6 a apple"
	var watermallan bool = true
	if watermallan {
		fruit = "1 apple"
	}
	fmt.Println(fruit)
	var money int = 11
	var busy bool = true
	switch {
	case money < 100:
		fmt.Println("money < 100")
		if busy {
			fmt.Println("busy")
			break
		}
		fmt.Println("not busy")
		fallthrough // 状态机 不管条件是否成立，都会执行下一个case
	case money == 100:
		fmt.Println("money == 100")
	default:
		fmt.Println("money > 100")
	}
}
