package main

import "fmt"

func main() {
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 =", f1)

	//自动推导类型
	//默认是float64，比float32更准确
	f2 := 3.14
	fmt.Printf("f2类型%T\n", f2)

}
