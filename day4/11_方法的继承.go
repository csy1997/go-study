package main

import "fmt"

type Person8 struct {
	name string
	sex  byte
	age  int
}

func (p *Person8) PrintInfo() {
	fmt.Printf("name = %v, sex = %c, age = %v\n", p.name, p.sex, p.age)
}

//继承Person8类，成员和方法都继承了
type Student8 struct {
	Person8
	id int
	addr string
}

func main() {
	s := Student8{Person8{"mike", 'm', 18}, 666, "beijing"}
	s.PrintInfo()
}