package main

import "fmt"

type Person5 struct {
	name string
	sex  byte
	age  int
}

//待接收者的函数叫方法
func (p Person5) PrintInfo() {
	fmt.Println("p =", p)
}

type char byte
//只要接收者不一样，方法重名也可以
func (c char) PrintInfo() {
	fmt.Println("c =", c)
}

type Int *int
//指针类型不能当函数接收者
//func (i Int) PrintInfo() {
//	fmt.Println("i =", i)
//}

//通过一个函数，给成员赋值
func (p *Person5) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	p := Person5{"mike", 'm', 18}
	p.PrintInfo()

	var p2 Person5
	(&p2).SetInfo("yoyo", 'f', 16)
	p2.PrintInfo()
}