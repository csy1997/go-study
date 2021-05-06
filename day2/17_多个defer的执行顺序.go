package main

import "fmt"

func main01() {
	//defer修饰的语句遵循先进后出
	defer fmt.Println("aaaaa")
	defer fmt.Println("bbbbb")

	//调用一个函数，导致内存出问题
	tt(0)//发生错误，前面的defer倒序执行

	defer fmt.Println("ccccc")
}

func main() {
	//先进后出
	defer fmt.Println("aaaaa")
	defer fmt.Println("bbbbb")

	//调用一个函数，导致内存出问题
	defer tt(0)//defer修饰的发生错误，后面的仍会倒序执行完

	defer fmt.Println("ccccc")
}

func tt(x int) {
	fmt.Println(100 / x)
}
