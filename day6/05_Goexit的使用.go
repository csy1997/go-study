package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("aaaaaaa")

		test()

		fmt.Println("ddddddd")
	}()

	//死循环不让主线程结束
	for {

	}
}

func test() {
	defer fmt.Println("ccccccc")

	//return//终止
	runtime.Goexit()//终止所在的协程

	fmt.Println("bbbbbbb")
}
