package main

import "fmt"

type Person2 struct {
	name string
	sex  byte
	age  int
}

type Student2 struct {
	Person2//只有类型，没有名字代表匿名字段，继承了Person类的成员
	id   int
	addr string
	name string//和Person2中的name同名
}

func main() {
	var s Student2
	s.name = "mike"
	//默认是就近原则，如果本作用与能找到name，就操作该成员
	//没找到就找继承的字段
	fmt.Printf("s = %+v\n", s)

	//显示调用
	s.Person2.name = "yoyo"
	fmt.Printf("s = %+v\n", s)
}