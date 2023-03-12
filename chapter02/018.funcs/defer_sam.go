package main

import (
	"fmt"
	"os"
	"time"
)

func deferGuess() {
	startTime := time.Now()
	fmt.Println("start time: ", startTime)
	// 约等于1-2百纳秒,很小的差距，因为代码走到这里括号里的代码已经执行了
	defer fmt.Println("duration 1:", time.Now().Sub(startTime))
	// 约等于1-2百纳秒,很小的差距，因为代码走到这里括号里的代码已经执行了
	defer printTime(time.Now().Sub(startTime))
	// 约等于5s
	defer printDiffTime(startTime)
	// 约等于5s
	defer func() {
		fmt.Println("duration upgraded:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("finish time: ", time.Now())
}

func printTime(time time.Duration) {
	fmt.Println("duration 2:", time)
}

func printDiffTime(startTime time.Time) {
	fmt.Println("duration 3:", time.Now().Sub(startTime))
}

func openFile() {
	fileName := "/"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer fileObj.Close()
}
