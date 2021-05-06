package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	modify(arr)
	fmt.Println(arr)
}

//数组作函数参数是值传递
//实参数组每个元素都要拷到形参数组里
func modify(arr [5]int) {
	arr[0] = 666
	fmt.Println(arr)
}
