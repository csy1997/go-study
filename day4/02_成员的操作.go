package main

import "fmt"

type Person1 struct {
	name string
	sex  byte
	age  int
}

type Student1 struct {
	Person1//只有类型，没有名字代表匿名字段，继承了Person类的成员
	id   int
	addr string
}

func main() {
	s1 := Student1{}
	s1.id = 1
	s1.addr = "beijing"
	s1.name = "mike"
	s1.sex = 'm'
	s1.age = 18

	fmt.Printf("s1 = %+v\n", s1)

	s1.Person1 = Person1{"jane", 'f', 16}
	fmt.Printf("s1 = %+v\n", s1)
}
