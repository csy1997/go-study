package main

import "fmt"

func main() {
	testa()
	testb()
	testc()
}

func testa() {
	fmt.Println("aaaaaaaaaaa")
}

func testb() {
	//fmt.Println("bbbbbbbbbbb")
	panic("this is a panic test")
}

func testc() {
	fmt.Println("ccccccccccc")
}
