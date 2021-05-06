package main

import "fmt"

func main() {
	MyFunc01(555)
	MyFunc02(111, 222)
}

func MyFunc01(a int) {
	//b = 111
	fmt.Println("b =", a)
}

func MyFunc02(/*b int, b int*/a, b int) {
	fmt.Printf("b = %d, b = %d\n", a, b)
}
