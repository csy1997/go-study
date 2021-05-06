package main

import (
	"fmt"
)

//全局变量，创建一个channel
var ch = make(chan int)

//定义一个打印机，参数为字符串，按每个字符打印，属于公共资源
func Printer1(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		//time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//让person3执行完了才能执行person4
func person3() {
	Printer1("hello")
	ch <- 666//给管道添加数据
}

func person4() {
	<- ch//从管道取数据，没有数据就会阻塞
	Printer1("world")
}

func main() {
	//两个人同时使用打印机
	go person3()
	go person4()

	//不让主协程结束
	for {

	}
}