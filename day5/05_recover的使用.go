package main

import "fmt"

func main() {
	testa2()
	testb2(10)
	testc2()
	testb2(1)//没有panic不打印
}

func testa2() {
	fmt.Println("aaaaaaaaaaa")
}

func testb2(x int) {
	//设置recover
	defer func() {
		//recover()
		//可以打印panic的错误信息
		//fmt.Println(recover())

		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 1
}

func testc2() {
	fmt.Println("ccccccccccc")
}