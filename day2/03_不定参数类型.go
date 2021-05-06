package main

import "fmt"

func main() {
	MyFunc03(1, 2)
	MyFunc04(1,2,3,4)
	MyFunc04()
}

func MyFunc03(a, b int)  {//固定参数
	
}

//...type不定参数类型
//注意：不定参数只能放在形参中最后一个参数，因为不定参数传参数量不固定，避免混淆
//func MyFunc04(args ...int, x int)  {
//
//}
func MyFunc04(args ...int) {
	fmt.Println("len(args) =", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}
}