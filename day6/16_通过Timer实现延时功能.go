package main

import (
	"fmt"
	"time"
)

func main() {
	//延时2s打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")

	<-time.After(2 * time.Second)//定时2秒，2秒后产生一个时间往channel写内容
	fmt.Println("时间又到了")
}
