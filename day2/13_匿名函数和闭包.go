package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	//匿名函数，无函数名
	f1 := func() {
		fmt.Println("b =", a)
		fmt.Println("str =", str)
	}
	f1()

	//定义匿名函数，同时调用（匿名函数必须要用）
	func() {
		fmt.Println("b =", a)
		fmt.Println("str =", str)
	}()

	//带参数的匿名函数
	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(1, 2)

	//匿名函数，有参有返回值
	max, min := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(1, 3)
	fmt.Printf("max = %d, min = %d\n", max, min)
}
