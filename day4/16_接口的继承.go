package main

import "fmt"

type Humaner13 interface {
	//方法只有声明，没有实现
	sayhi()
}

type Personer13 interface {
	Humaner13
	sing(lrc string)
}

type Student13 struct {
	name string
	id int
}

func (s *Student13) sayhi() {
	fmt.Printf("Student %v, %v sayhi\n", s.id, s.name)
}

func (s *Student13) sing(lrc string) {
	fmt.Printf("Student %v 在唱着 %v\n", s, lrc)
}

func main() {
	var i Personer13
	s := &Student13{"mike", 666}
	i = s

	i.sayhi()//继承过来的方法
	i.sing("sdfeq")
}