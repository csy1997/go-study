package main

import "fmt"

//func main() {
//	fmt.Println(test11())
//	fmt.Println(test11())
//	fmt.Println(test11())
//	fmt.Println(test11())
//}

func main() {
	//返回值为一个匿名函数，赋给f，通过f来调用该闭包函数
	//不关心这些捕获了的变量和常量是否已经超出了作用域
	//只要有闭包还在使用，这些变量就会一直存在
	f := test12()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //25
}

func test11() int {
	//函数被调用时，x才分配空间，才初始化为0
	var x int
	x++
	return x * x//函数调用完毕，x自动释放
}

//函数的返回值是一个匿名函数，返回一个函数类型
func test12() func() int {
	var x int//没有初始化，值为0
	return func() int {
		x++
		return x * x
	}
}
