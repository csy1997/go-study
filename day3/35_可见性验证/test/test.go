package test

import "fmt"

type student struct {
	id int
	Name string
}

type Student struct {
	id int
	Name string
}

func myFunc() {
	fmt.Println("test")
}

func MyFunc() {
	fmt.Println("test")
}