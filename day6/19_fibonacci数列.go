package main

import "fmt"

//生产者放数，消费者取数，一次只能一个，消费者决定是否停止
func main() {
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //程序是否结束

	//消费者，从channel中读取内容
	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch//取数
			fmt.Println("num =", num)
		}
		quit <- true//取完10个后停
	}()

	//生产者，产生数字，写入channel
	fibonacci(ch, quit)
}

//函数里ch只写，quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {//监听ch和quit（哪个可以进行执行哪个，都可以就随机一个）
		case ch <- x://初始为空或外面消费者取了数，就可以放新数了
			x, y = y, x+y
		case flag := <-quit://外面消费者放入停止标识了，就取出来
			fmt.Println("flag =", flag)
			//break//break不是跳出循环，会再次进入for，消费者执行完了，会死锁
			return
		}
	}
}
