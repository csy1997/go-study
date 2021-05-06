package main

import "fmt"

func main() {
	var ch byte = 97
	fmt.Println("ch =", ch)
	//格式化输出，%c以字符方式打印，%d以整形方式打印
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'a'
	fmt.Printf("%c, %d\n", ch, ch)

	//大小写转换，大比小少32
	fmt.Printf("%c\n", 'A'+32)
	fmt.Printf("%c\n", 'a'-32)

	//以"/"开头的是转义字符
	fmt.Printf("hello go%c", '\n')
	fmt.Printf("hello itcast")
}
