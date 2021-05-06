package main

import "fmt"

func main() {
	a := Calc(1, 2, add1)
	fmt.Println("b =", a)
}

type FuncType1 func(int, int) int

func add1(a, b int) int {
	return a + b
}

//函数有一个参数是函数类型，就是回调函数
//比如计算器实现四则运算
//多态的一种实现，调用同一个接口，实现不同的表现
//现有想法，后面再实现功能
func Calc(a, b int, fTest FuncType1) (res int) {
	fmt.Println("Calc")
	res = fTest(a, b)
	return
}
