package main

import "fmt"

type Student5 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	s1 := Student5{1, "mike", 'm', 18, "beijing"}
	s2 := Student5{1, "mike", 'm', 18, "beijing"}
	s3 := Student5{2, "mike", 'm', 18, "beijing"}

	fmt.Println("s1 == s2:", s1 == s2)
	fmt.Println("s1 == s3:", s1 == s3)

	//同类型的两个结构体可以相互赋值
	var tmp Student5
	tmp = s3
	fmt.Println("tmp =", tmp)
}