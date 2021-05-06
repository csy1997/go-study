package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个有缓存的channel
	ch := make(chan int, 3)

	//len剩余数据个数，cap缓冲区大小，无缓冲的channel都是0
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println("子协程：i =", i)
			fmt.Printf("放num = %d, len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
			ch <- i//放入之后没有取，下次再放会阻塞
		}
	}()

	//等2秒
	time.Sleep(2*time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("取num =", num)
	}
}