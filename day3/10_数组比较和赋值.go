package main

import "fmt"

func main() {
	//支持比较，只能==和!=，类型必须相等，然后判断是否每个元素一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	//d := [6]int{1, 2, 3, 4, 5, 6}
	//e := [5]string{"1","2","3","4","5"}
	//f := [5]byte{'1','2','3','4','5'}

	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c)
	//fmt.Println("a == d", a == d)
	//fmt.Println("a == e", a == e)
	//fmt.Println("a == f", a == f)

	//同类型的数组可以赋值
	var g [5]int
	g = a
	fmt.Println("g =", g)
}
