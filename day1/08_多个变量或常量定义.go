package main

import "fmt"

func main() {
	//var a int
	//var b float64

	var (
		a int     = 1
		b float64 = 3.14
	)
	const (
		i = 1
		j = 3.14
	)

	fmt.Printf("a = %d, b = %f, i = %d, j = %f", a, b, i, j)
}
