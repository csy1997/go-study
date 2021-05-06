package main

import "fmt"

func main() {
	testa1()
	testb1(10)
	testc1()
}

func testa1() {
	fmt.Println("aaaaaaaaaaa")
}

func testb1(x int) {
	//fmt.Println("bbbbbbbbbbb")
	//panic("this is a panic test")
	var a [10]int
	a[x] = 1//当x为10的时候，导致数组越界，产生一个panic，导致程序崩溃
}

func testc1() {
	fmt.Println("ccccccccccc")
}

