package main

import "fmt"

func main() {
	//数组[]里面的长度是固定的一个常量，不能修改长度
	a := [5]int{}
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	//切片[]里为空或者...，长度和容量可以不固定
	s := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	//末尾增加一个成员
	s = append(s, 11)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	//自动推导类型，同时初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)

	//借助make函数，格式为make(切片类型, 长度, 容量)
	s2 := make([]int, 5, 10)
	fmt.Println(len(s2), ",", cap(s2))

	//没有指定容量，和长度一样
	s3 := make([]int, 5)
	fmt.Println(len(s3), ",", cap(s3))
}
