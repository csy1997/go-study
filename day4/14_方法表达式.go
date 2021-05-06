package main

import "fmt"

type Person11 struct {
	name string
	sex  byte
	age  int
}

func (p Person11) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person11) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, *p)
}

func main() {
	p := Person11{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	p.SetInfoPointer()	//传统调用方式

	//方法表达式
	f := (*Person11).SetInfoPointer
	f(&p)//显式地把接收者传递过去，等价于p.SetInfoPointer()
}