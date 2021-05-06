package main

import "fmt"

type Humaner12 interface {
	//方法只有声明，没有实现
	sayhi()
}

type Student12 struct {
	name string
	id int
}

func (s *Student12) sayhi()  {
	fmt.Printf("Student %v, %v sayhi\n", s.id, s.name)
}

type Teacher12 struct {
	addr string
	group string
}

func (t *Teacher12) sayhi()  {
	fmt.Printf("Teacher %v, %v sayhi\n", t.addr, t.group)
}

type MyStr string

func (m *MyStr) sayhi() {
	fmt.Printf("MyStr %v sayhi\n", m)
}

//函数参数为借口类型，实现了多态，同一个函数可以有不同表现
func WhoSayHi(i Humaner12) {
	i.sayhi()
}

func main() {
	//定义接口类型的变量
	var i Humaner12

	//只要实现了此接口方法的类型，那么这个类型的指针变量（指接收者类型）就可以给i赋值
	s := &Student12{"mike", 666}
	i = s
	i.sayhi()

	t := &Teacher12{"beijing", "go"}
	i = t
	i.sayhi()

	var str MyStr = "hello mike"
	i = &str
	i.sayhi()

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner12, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, i := range x {
		WhoSayHi(i)
	}
}