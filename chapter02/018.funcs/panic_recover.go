package main

import (
	"fmt"
)

func panicAndRecover() {
	defer coverPanicUpgraded("1") // 这样是能抓住严重错误的

	defer func() { // 抓不住严重错误
		// 多一层func，抓不住上层严重错误
		defer coverPanicUpgraded("2")
		// 能抓住这里的错误
		errorFn()
	}()

	// 抓不住严重错误
	defer coverPanic()
	errorFn()
}

func errorFn() {
	var nameScore map[string]int = nil
	nameScore["进文"] = 100
}

func coverPanic() { // 未能抓住panic
	// recover() // 不能单独使用recover
	//  cover不住上层的panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出严重故障：", 3, r)
		}
	}()
}

func coverPanicUpgraded(site string) {
	if r := recover(); r != nil {
		fmt.Println("系统出严重故障：", site, r)
	}
}
