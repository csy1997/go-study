package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片，让别的协程先执行，执行完后再回来执行此协程
		runtime.Gosched()
		fmt.Println("hello")
	}
}
