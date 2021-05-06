package main

import "fmt"

func main() {
	result := add(1, 1)
	fmt.Println("result =", result)

	//函数类型变量声明与赋值
	var fTest FuncType
	fTest = add
	result = fTest(10, 20)//和add功能一样
	fmt.Println("result =", result)
	fTest = minus
	result = fTest(10, 5)//和minus功能一样
	fmt.Println("result =", result)
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数类型起名
//FuncType是一个函数类型，无函数名，无{}
type FuncType func(int, int) int