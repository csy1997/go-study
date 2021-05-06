package main

import "fmt"

func main() {
	var a bool
	//没有初始化默认为false
	fmt.Println("a =", a)
	a = true
	fmt.Println("a =", a)

	//自动推导类型
	var b = false
	fmt.Println("b =", b)

	c := true
	fmt.Println("c =", c)
}
