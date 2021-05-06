package main

import "fmt"

type Student3 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	var s Student3

	//用"."操作成员
	s.id = 1
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.address = "beijing"
	fmt.Println(s)
}
