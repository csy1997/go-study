package main

import "fmt"

func main() {
	a := 10
	str := "abc"

	func() {
		//闭包是以引用方式捕获外部变量的
		a = 666
		str = "go"
		fmt.Printf("内部：b = %d, str = %s\n", a, str)
	}()

	fmt.Printf("外部：b = %d, str = %s\n", a, str)
}


