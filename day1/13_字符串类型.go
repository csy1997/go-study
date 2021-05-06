package main

import "fmt"

func main() {
	var str1 string
	str1 = "abc"
	fmt.Println("str1 =", str1)

	//自动推导
	str2 := "mike"
	fmt.Printf("str2的类型是%T\n", str2)

	//内建函数，len()可以返回字符串长度
	fmt.Println("str2长度：", len(str2))
}
