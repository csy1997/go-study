package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for { //for后面不写任何东西，死循环
		i++
		time.Sleep(time.Second)

		if i == 5 {
			break
			//contiue//跳到下一次循环
		}
		fmt.Println("i =", i)
	}
}