package main

import "fmt"

type Humaner14 interface {//子集
	//方法只有声明，没有实现
	sayhi()
}

type Personer14 interface {//超集
	Humaner14
	sing(lrc string)
}

type Student14 struct {
	name string
	id int
}

func (s *Student14) sayhi() {
	fmt.Printf("Student %v, %v sayhi\n", s.id, s.name)
}

func (s *Student14) sing(lrc string) {
	fmt.Printf("Student %v 在唱着 %v\n", s, lrc)
}

func main() {
	//超集可以转化为子集，反之不行
	var iPro Personer14
	var i Humaner14

	//iPro = i//子集没有sing方法，不能转为iPro
	s := &Student14{"mike", 666}
	iPro = s
	i = iPro

	i.sayhi()
}
