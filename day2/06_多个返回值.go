package main

import "fmt"

func main() {
	a, b, c := myFunc09()
	fmt.Printf("b = %d, b = %d, c = %d\n", a, b, c)

	a,b,c = myFunc10()
	fmt.Printf("b = %d, b = %d, c = %d\n", a, b, c)
}

func myFunc09() (int, int, int) {
	return 1, 2, 3
}

//常用写法
func myFunc10() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}
