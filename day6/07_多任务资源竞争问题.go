package main

import (
	"fmt"
)

//定义一个打印机，参数为字符串，按每个字符打印，属于公共资源
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		//time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
}

func person2() {
	Printer("world")
}

func main() {
	//两个人同时使用打印机
	go person1()
	go person2()

	//不让主协程结束
	for {

	}
}
