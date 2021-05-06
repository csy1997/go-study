package main

import "fmt"

func main() {
	var a int
	a = myFunc06()
	fmt.Println("b =", a)

	b := myFunc06()
	fmt.Println("b =", b)

	c := myFunc07()
	fmt.Println("c =", c)

	d := myFunc08()
	fmt.Println("d =", d)
}

func myFunc06() int {
	return 666
}

//给返回值起一个变量名，go语言推荐写法
func myFunc07() (res int) {
	return 666
}

//最常用写法
func myFunc08() (res int) {
	res = 666
	return
}
