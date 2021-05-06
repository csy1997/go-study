package main

import "fmt"

func main() {
	var a int
	//每个变量有两层含义：变量的内存，变量的地址
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %v\n", &a)

	//用指针类型保存变量的地址，*int保存int的地址，**int保存*int的地址
	var p *int
	p = &a//指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p = %v, &a = %v\n", p, &a)

	*p = 666//*p操作的不是p的内存，是p所指向的内存（即a）
	fmt.Printf("*p = %v, a = %v\n", *p, a)

	var p1 *int
	fmt.Println("p1 =", p1)

	//err，p1没有合法指向，不能操作没有合法指向的内存
	*p1 = 666
	fmt.Println("*p1 =", *p1)
}