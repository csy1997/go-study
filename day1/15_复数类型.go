package main

import "fmt"

func main() {
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t =", t)

	//自动推导
	t2 := 3.3 + 4.4i
	fmt.Printf("t2类型是%T\n", t2)

	//通过内建函数，取实部和虚部
	fmt.Printf("real(t2) = %f, imag(t2) = %f\n", real(t2), imag(t2))
}
