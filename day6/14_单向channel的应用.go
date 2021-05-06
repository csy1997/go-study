package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	//生产者，生产数据写入channel
	//channel传参是引用传递
	go producer(ch)

	//消费者，读取数据并打印
	consumer(ch)
}

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		out <- i*i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num =", num)
	}
}
