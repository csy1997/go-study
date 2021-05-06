package main

import "fmt"

type mystr string

type Person3 struct {
	name string
	sex  byte
	age  int
}

type Student3 struct {
	Person3//只有类型，没有名字代表匿名字段，继承了Person类的成员
	int //基础类型的匿名字段
	mystr //自定义类型的匿名字段
}

func main() {
	s := Student3{Person3{"mike", 'm', 18}, 1, "beijing"}
	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
	fmt.Println(s.Person3, s.int, s.mystr)
}