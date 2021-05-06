package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[0:3:5]
	fmt.Println(s)
	fmt.Println("len(s) =", len(s))//3-0
	fmt.Println("cap(s) =", cap(s))//5-0
}
