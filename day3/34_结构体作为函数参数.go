package main

import "fmt"

type Student6 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	s := Student6{1, "mike", 'm', 18, "beijing"}
	test01(s)//值传递，实参不改变
	fmt.Println(s)

	test02(&s)//地址传递
	fmt.Println(s)
}

func test01(s Student6) {
	s.id = 666
	fmt.Println(s)
}

func test02(p *Student6) {
	p.id = 666
	fmt.Println(*p)
}