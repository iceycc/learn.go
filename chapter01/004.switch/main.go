package main

import (
	"fmt"
)

func main() {
	var money int = 20
	var busy bool = true

	switch money {
	case 20:
		fmt.Println("点个外卖")
		fallthrough // 状态机 不管条件是否成立，都会执行下一个case
	case 200:
		fmt.Println("下个馆子")
		if busy {
			break
		} else {
			fmt.Println("再吃点零食")
		}
		// ....
	}
	fmt.Println("end")
}
