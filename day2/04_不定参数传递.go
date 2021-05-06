package main

import "fmt"

func main() {
	test(1, 2, 3, 4)
}

func test(args ...int) {
	//全部元素传给myFunc05
	myFunc05(args...)
	fmt.Println()
	//把下标为2之后（包括2）的元素传递给myFunc05使用
	myFunc05(args[2:]...)
	fmt.Println()
	//把下标为2之前（不包括2）的元素传递给myFunc05使用
	myFunc05(args[:2]...)
}

func myFunc05(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data =", data)
	}
}
