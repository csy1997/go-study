package main

import "fmt"

func main() {
	//[n]，n为数组元素个数
	var a [10]int
	var b [5]int
	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))

	//定义数组指定长度n必须是常量
	//n := 10
	//var c [n]int

	//操作数组元素，从0开始到len()-1，位置叫下标
	a[0] = 1
	i := 2
	a[i] = 2

	//赋值
	for i := 0; i < len(a); i++ {
		a[i] = i+1
	}

	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}
}
