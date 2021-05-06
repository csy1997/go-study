package main

import (
	"fmt"
	"strings"
)

func main() {
	//Contains
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "hello"))

	//Join
	s1 := []string{"abc", "defg", "hi", "jkl"}
	buf1 := strings.Join(s1, "~")
	fmt.Println(buf1)
	buf2 := strings.Join(s1, "")
	fmt.Println(buf2)

	//Index
	fmt.Println(strings.Index("hellogo", "go"))
	fmt.Println(strings.Index("hellogo", "abc"))//不存在返回-1

	//Repeat
	buf3 := strings.Repeat("go", 3)
	fmt.Println(buf3)

	//Replace，将old子串替换为new子串，n为从前到后替换个数，负数为全替换
	fmt.Println(strings.Replace("hello go", "o", "~", 1))
	fmt.Println(strings.Replace("hello go", "o", "m", 10))
	fmt.Println(strings.Replace("hello go", "o", "m", -1))
	fmt.Println(strings.Replace("hello go", "god", "", -1))

	//Split，以指定的分隔符拆分
	buf4 := "abc~defg~hi~jkl"
	s2 := strings.Split(buf4, "~")
	fmt.Println(s2)

	//Trim去掉两头的字符
	buf5 := strings.Trim("    are u ok    ", " ")
	fmt.Println(buf5)

	//Fields，去掉所有空格，把分隔得到的元素放入切片
	s3 := strings.Fields("    are u ok    ")
	for i, data := range s3 {
		fmt.Println(i, ":", data)
	}
}
