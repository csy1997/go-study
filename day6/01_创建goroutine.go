package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a new task")
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask() //创建一个协程，新建一个任务

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}

}
