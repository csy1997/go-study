package main

import "fmt"

func main() {
	//goto可以用在任何地方，但不能跨函数使用
	fmt.Println("11111")
	goto END

	fmt.Println("22222")
	END:
		fmt.Println("33333")
}
