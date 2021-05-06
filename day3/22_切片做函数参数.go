package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 10

	//创建一个切片，len为n
	s := make([]int, n)
	InitData(s)
	fmt.Println("s =", s)

	BubbleSort(s)
	//切片做函数参数是引用传递，函数内修改会影响实参
	fmt.Println("s =", s)
}

func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func InitData(s []int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}
