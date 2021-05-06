package main

import "fmt"

type Student struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	//顺序初始化，每个成员都必须初始化
	//var s1 Student = Student{1}
	var s1 Student = Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println(s1)

	//指定成员初始化
	var s2 Student = Student{name: "mike", address: "beijing"}
	fmt.Println(s2)
}