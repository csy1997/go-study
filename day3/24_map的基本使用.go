package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println("m1 =", m1)
	fmt.Println("len(m1) =", len(m1))//map只有len

	//通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 =", m2)
	fmt.Println("len(m2) =", len(m2))

	//指定长度，但初始没有数据
	m3 := make(map[int]string, 2)
	fmt.Println("m3 =", m3)
	fmt.Println("len(m3) =", len(m3))

	m4 := make(map[int]string, 2)
	fmt.Println("len(m4) =", len(m4))
	//添加元素超过长度会自动扩容，map无序
	m4[0] = "mike"
	m4[-1] = "go"
	m4[6] = "c++"
	fmt.Println("m4 =", m4)
	fmt.Println("len(m4) =", len(m4))

	//初始化，键值唯一
	m5 := map[int]string{1: "mike", 2: "go", 3: "c++"}
	fmt.Println("m5 =", m5)
	fmt.Println("len(m5) =", len(m5))
}
