package main

import "fmt"

func main() {
	//a := 10
	//b := 20
	//c := 30

	a, b := 10, 20
	//交换两个变量值
	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)

	//_匿名变量，丢弃变量不处理，配合函数使用
	temp, _ := a, b
	fmt.Println("temp = ", temp)

	var d, e int
	_, d, e = test()
	fmt.Printf("d = %d, e = %d\n", d, e)
}

func test() (a, b, c int) {
	return 1, 2, 3
}
