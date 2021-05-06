package main

import "fmt"

//定义在函数外部的变量是全局变量
//全局变量在任何地方都能使用
var b int

func main() {
	b = 10
	fmt.Println("b =", b)
	test()
}

func test13() {
	fmt.Println("b =", b)
}