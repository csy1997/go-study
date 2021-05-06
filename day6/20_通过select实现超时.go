package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num =", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时！")
				quit <- true
			}
		}
	}()

	//写入数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			x := rand.Intn(4) + 1
			time.Sleep(time.Duration(x) * time.Second)
			fmt.Printf("等待%d秒\n", x)
		}
	}()

	<-quit
	fmt.Println("超时程序结束")
}
