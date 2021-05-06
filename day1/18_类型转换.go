package main

import "fmt"

func main() {
	var flag bool = true
	fmt.Printf("flag = %t\n", flag)
	//bool不能转换为int，int也不能转为bool
	//fmt.Printf("flag = %d\n", int(flag))
	//flag = bool(1)

	var ch byte
	ch = 'a'
	var t int
	//t = ch
	t = int(ch) // 用int强转
	fmt.Printf("t = %d", t)
}
