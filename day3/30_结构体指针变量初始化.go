package main

import "fmt"

type Student2 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	//顺序初始化，每个成员都必须初始化，别忘&
	var p1 *Student2 = &Student2{1, "mike", 'm', 18, "beijing"}
	fmt.Println(*p1)

	//指定成员初始化
	p2 := &Student2{name: "mike", address: "beijing"}
	fmt.Printf("p2的类型：%T\n", p2)
	fmt.Println(p2)
}
