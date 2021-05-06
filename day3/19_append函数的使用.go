package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	//在原切片末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 =", s1)
}
