package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	modify2(&arr)//地址传递
	fmt.Println(arr)
}

//p指向数组arr，*p代表指针指向的内存，就是实参a
func modify2(p *[5]int) {
	(*p)[0] = 666
	fmt.Println(*p)
}
