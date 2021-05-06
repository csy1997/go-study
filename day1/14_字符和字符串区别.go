package main

import "fmt"

func main() {
	var ch byte
	var str string
	//1.单双引号
	//2.字符往往只有一个字符，转义字符除外；字符串可以一个可以多个
	//3.字符串都是隐藏了一个结束符'/0'
	ch = 'a'
	fmt.Println("ch =", ch)
	str = "a"

	//只想操作字符串的某个字符
	str = "hello go"
	fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])
}
