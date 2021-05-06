package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for true {
			i++
			fmt.Println("子协程 i =", i)
			time.Sleep(time.Second)
		}
	}()
}