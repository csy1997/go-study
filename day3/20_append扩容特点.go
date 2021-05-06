package main

import "fmt"

func main() {
	//如果超过原来的容量，通常扩容为2倍
	s := make([]int, 0, 1)
	fmt.Println("cap =", cap(s))
	for i := 0; i < 8; i++ {
		s = append(s, i)
		fmt.Println("cap =", cap(s))
	}
}
