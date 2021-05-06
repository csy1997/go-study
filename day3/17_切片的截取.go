package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max]取下标从low开始的元素，len=high-low，cap=max-low
	s1 := arr[:]//不指定容量则和长度一样
	fmt.Println(s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	//操作某个元素，和数组操作方式一样
	data1 := arr[1]
	fmt.Println(data1)

	//下标3~5，容量为4
	s2 := arr[3:6:7]
	fmt.Println(s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	//0~5，长度为6
	s3 := arr[:6]
	fmt.Println(s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	//3~9，长度为7
	s4 := arr[3:]
	fmt.Println(s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))
}