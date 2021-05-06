package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//新切片
	s1 := arr[2:5]//容量为8
	s1[1] = 666
	fmt.Println("s1 =", s1)
	fmt.Println("arr =", arr)

	//对切片切片
	s2 := s1[2:7]
	s2[2] = 777
	fmt.Println("s2 =", s2)
	fmt.Println("arr =", arr)
}
