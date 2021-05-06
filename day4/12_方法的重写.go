package main

import "fmt"

type Person9 struct {
	name string
	sex  byte
	age  int
}

func (p *Person9) PrintInfo() {
	fmt.Printf("name = %v, sex = %c, age = %v\n", p.name, p.sex, p.age)
}

type Student9 struct {
	Person9
	id int
	addr string
}

//和Person9的方法同名，叫重写
func (s *Student9) PrintInfo() {
	fmt.Println("s =", s)
}

func main() {
	s := Student9{Person9{"mike", 'm', 18}, 666, "beijing"}
	//就近原则，先找本作用域的方法，找不到再找继承的方法
	s.PrintInfo()

	//用Person9的方法直接显示调用
	s.Person9.PrintInfo()
}