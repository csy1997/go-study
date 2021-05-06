package main

import "fmt"

type Person7 struct {
	name string
	sex  byte
	age  int
}

func (p Person7) ValueTest() {
	fmt.Println("ValueTest")
}

func (p *Person7) PointerTest() {
	fmt.Println("PointerTest")
}

func main() {
	//结构体作为指针变量(*Person7)和普通变量，能调用的所有方法叫方法集
	p1 := &Person7{"mike", 'm', 18}
	//内部根据需要会将p1和*p1相互转换，两种方法都能调用
	p1.PointerTest()
	(*p1).PointerTest()
	(*p1).ValueTest()
	p1.ValueTest()

	//内部根据需要会将p2和&p2相互转换，两种方法都能调用
	p2 := Person7{"mike", 'm', 18}
	p2.ValueTest()
	(&p2).ValueTest()
	(&p2).PointerTest()
	p2.PointerTest()
}