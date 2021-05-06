package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "go", 3: "yoyo"}

	//第一个值key，第二个value，遍历无序
	for key, value := range m {
		fmt.Printf("%d:%s\n", key, value)
	}

	//判断一个key是否存在，第一个为key对应的value，第二个为是否存在
	value, ok := m[0]
	fmt.Printf("%v, %v\n", value, ok)
	value, ok = m[1]
	fmt.Printf("%v, %v\n", value, ok)
}
