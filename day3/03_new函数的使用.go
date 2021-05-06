package main

import "fmt"

func main() {
	a := 10
	var p *int

	//p = &a
	p = new(int)
	*p = a
	fmt.Println("*p =", *p)

	q := new(int)
	*q = 777
	fmt.Println("q =", *q)
}
