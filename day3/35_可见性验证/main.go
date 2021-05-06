package main

import (
	"fmt"
	"study/day3/35_可见性验证/test"
)

func main() {
	//小写函数，不再同一个包，不可用
	//test.myFunc()
	test.MyFunc()

	//结构体同理
	//fmt.Println(test.student{})
	s := test.Student{}
	fmt.Println(s)

	//结构体中的成员同理
	//s.id = 1
	s.Name = "Jane"
	fmt.Println(s)
}

