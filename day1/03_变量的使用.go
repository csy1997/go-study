package main

import "fmt" //导入包必须要使用

func main() {
	//变量声明必须要用
	//只声明不初始化的变量也有初始值
	//同一个{}中声明的变量名是唯一的
	var a int
	fmt.Println("a = ", a)

	a = 10//变量赋值
	fmt.Println("a = ", a)

	var b int = 15//声明同时赋值
	fmt.Println("b = ", b)

	//自动推导类型
	c := 30
	//格式化输出类型
	fmt.Printf("c的类型是 %T", c)
}
