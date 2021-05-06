package main

import "fmt"

//可接收任意类型和个数的参数
func xxx(arg ...interface{})  {
	fmt.Println("xxx")
}

func main() {
	//空接口万能类型，可以保存任意类型的值
	var i interface{} = 1
	fmt.Println(i)
	i = "abc"
	fmt.Println(i)
	i = true
	fmt.Println(i)

	xxx(1, "abc", true)
}
