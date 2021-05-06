package main

import "fmt"

type Student4 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	var s Student4

	var p *Student4
	p = &s

	//通过指针操作成员，p.id和(*p).id完全等价
	fmt.Println(p.id)
	fmt.Println((*p).id)
	p.age = 18
	p.address = "beijing"
	fmt.Println(p)

	//通过new关键字申请一个结构体
	p2 := new(Student4)
	p2.age = 18
	p2.address = "beijing"
	fmt.Println(p2)
}
