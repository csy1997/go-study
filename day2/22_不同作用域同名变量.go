package main

import "fmt"

var a byte

func main() {
	var a int

	//不同作用域，允许定义不同类型同名变量
	//使用变量的原则是就近原则
	fmt.Printf("a的类型：%T\n", a)
	{
		var a float32
		fmt.Printf("a的类型：%T\n", a)
	}
	test14()
}

func test14() {
	fmt.Printf("a的类型：%T\n", a)
}