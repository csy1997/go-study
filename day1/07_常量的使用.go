package main

import "fmt"

func main() {
	//变量运行期间可以改变，用var
	//常量运行期间不能改变，用const
	const a int = 10
	//a = 20
	fmt.Println("a =", a)

	//const b := 20 //常量不能用:=
	const b = 20
	fmt.Printf("b的类型是%T\n", b)
	fmt.Println("b =", b)
}
