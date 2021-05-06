package main

import "fmt"

func main() {
	//for 初始化条件; 判断条件; 条件变化 {}

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)

	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	//迭代打印每个元素，默认返回2个值，一个元素位置，一个元素本身
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}
}
