package main

import "fmt"

type Person10 struct {
	name string
	sex  byte
	age  int
}

func (p Person10) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person10) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, *p)
}

func main() {
	p := Person10{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	p.SetInfoPointer()	//传统调用方式

	//保存入口方法地址
	pFunc := p.SetInfoPointer //这个就是方法值，调用函数时，无需再传接收者，隐藏了接收者
	pFunc()

	vFunc := p.SetInfoValue
	vFunc()
}