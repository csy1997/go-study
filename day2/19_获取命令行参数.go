package main

import (
	"fmt"
	"os"
)

func main() {//go run 19_获取命令行参数.go hello world
	//接收用户传递的参数，都是以字符串方式传递的
	list := os.Args

	n := len(list)
	fmt.Println("n =", n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}
}