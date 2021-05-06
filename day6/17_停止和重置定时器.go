package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	time.Sleep(time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到")
	}()

	timer.Reset(1 * time.Second)//重新将定时设置为1s（从当前时间开始）
	//timer.Stop()//停止定时器

	for {}
}
