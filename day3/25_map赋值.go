package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "yoyo"}
	fmt.Println(m1)
	m1[1] = "jane"//key已存在会替换value
	fmt.Println(m1)
	m1[3] = "go"//不存在会增加
	fmt.Println(m1)
}
