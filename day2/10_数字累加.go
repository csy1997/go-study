package main

import "fmt"

func main() {
	fmt.Println("sum =", test01())
	fmt.Println("sum =", test02(100))
}

func test01() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func test02(i int) (sum int) {
	if i == 1 {
		return 1
	}
	sum = i + test02(i-1)
	return
}
