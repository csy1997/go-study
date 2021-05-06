package main

import "fmt"

func main() {
	a, b := 10, 20

	//通过一个函数交换a和b的内容
	swap(a, b)//变量本身传递，值传递非地址传递
	fmt.Printf("a = %d, b = %d\n", a, b)

	swap1(&a, &b)//地址传递
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(a int, b int) {
	a, b = b, a
	//fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap1(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}
