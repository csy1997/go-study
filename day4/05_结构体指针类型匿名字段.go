package main

import "fmt"

type Person4 struct {
	name string
	sex  byte
	age  int
}

type Student4 struct {
	*Person4//只有类型，没有名字代表匿名字段，继承了Person类的成员
	id   int
	addr string
}

func main() {
	s := Student4{&Person4{"mike", 'm', 18}, 1, "beijing"}
	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.name, s.age, s.sex, s.id, s.addr)

	var s2 Student4
	s2.Person4 = new(Person4)
	s2.id = 1
	s2.addr = "beijing"
	s2.name = "mike"
	s2.sex = 'm'
	s2.age = 18
	fmt.Println(s2.name, s2.age, s2.sex, s2.id, s2.addr)
}