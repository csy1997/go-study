package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person//只有类型，没有名字代表匿名字段，继承了Person类的成员
	id   int
	addr string
}

func main() {
	var s1 = Student{Person{"mike", 'm', 18}, 1, "beijing"}
	fmt.Println(s1)

	s2 := Student{Person{"mike", 'm', 18}, 1, "beijing"}
	//%+v格式，显示更详细
	fmt.Printf("s2 = %+v\n", s2)

	//指定成员初始化，没初始化的为0值
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v\n", s3)
	s4 := Student{Person: Person{name: "mike"}, id: 1}
	fmt.Printf("s4 = %+v\n", s4)
}
