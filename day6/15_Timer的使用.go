package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置为2秒后，往time通道写内容（当前时间开始），且只会响应一次（多取会报错）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2s后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C
	fmt.Println("t =", t)
}
