package main

import "fmt"

func main() {

	//延迟调用，main函数结束前调用
	defer fmt.Println("bbbbbb")

	fmt.Println("aaaaaa")
}
