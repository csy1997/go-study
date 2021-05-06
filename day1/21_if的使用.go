package main

import "fmt"

func main() {
	s := "张三"
	if s == "张三" {
		fmt.Println("233")
	}

	//if支持1个初始化语句，初始化语句和判断语句用;隔开
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}

	a := 10
	if a == 10 {
		fmt.Println("= 10")
	} else if a > 10 {
		fmt.Println("> 10")
	} else if a < 10 {
		fmt.Println("< 10")
	} else {
		fmt.Println("不可能")
	}
}
