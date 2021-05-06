package main

import "fmt"

type Person6 struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值
//接收者为普通变量，值传递修改的是拷贝的值
func (p Person6) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue地址：%p\n", &p)
}

//接收者为指针变量，引用传递修改的是原始值
func (p *Person6) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoPoint地址：%p\n", p)
}

func main() {
	var s1 Person6
	fmt.Printf("s1地址：%p\n", &s1)

	//值语义，修改完不会影响外面
	s1.SetInfoValue("mike", 'm', 18)
	fmt.Println("s1 =", s1)

	//引用语义
	//(&s1).SetInfoPointer("mike", 'm', 18)//下同
	s1.SetInfoPointer("mike", 'm', 18)
	fmt.Println("s1 =", s1)
}