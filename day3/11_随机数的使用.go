package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子，只需一次
	//如果种子参数一样，每次得到的随机数都一样
	rand.Seed(666)
	//产生随机数
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int())
	}

	//以系统当前时间为参数
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int())
	}

	//限制在100以内，包括0不包括100
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}
}
