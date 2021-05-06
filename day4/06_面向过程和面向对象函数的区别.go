package main

import "fmt"

func Add01(a, b int) (res int) {
	res = a + b
	return
}

type long int

//面向对象方法，给某个类型绑定一个函数
func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	res1 := Add01(1, 1)
	fmt.Println("res1 =", res1)

	var a long = 2
	//变量名.函数(所需参数)
	res2 := a.Add02(1)
	fmt.Println("res2 =", res2)

	//面向对象只是换了一种表现形式
}
