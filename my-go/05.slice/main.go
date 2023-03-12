package main

import "fmt"

func main() {
	a := []int{}
	a = append(a, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	fmt.Println(a)
	fmt.Println(a[2:3])
	fmt.Println(a)

	b := []int{11, 22, 33, 44, 55, 66}
	b = append(b[:2], b[3:]...)
	fmt.Println(b)

}
