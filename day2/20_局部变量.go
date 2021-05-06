package main

import "fmt"

func main() {
	//定义在{}的变量是局部变量，只在{}里有效
	//执行到定义变量那句话才开始为其分配空间，离开作用域自动释放
	{
		i := 10
		fmt.Println("i =", i)
	}
	//i = 111

	if flag := 3; flag == 3 {
		fmt.Println("true")
	}
	//fmt.Println(flag)
}
