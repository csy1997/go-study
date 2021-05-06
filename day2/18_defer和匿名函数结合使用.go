package main

import "fmt"

func main02() {
	a := 10
	b := 20
	defer func() {//后执行
		fmt.Printf("b = %d, b = %d\n", a, b)
	}()

	a = 111
	b = 222
	fmt.Printf("外部：b = %d, b = %d\n", a, b)
}

func main() {
	a := 10
	b := 20
	defer func(a, b int) {//后执行
		fmt.Printf("b = %d, b = %d\n", a, b)
	}(a, b) //虽然加了defer，但参数还是先传过来了，只是没有调用，defer的是最后执行该函数

	a = 111
	b = 222
	fmt.Printf("外部：b = %d, b = %d\n", a, b)
}
