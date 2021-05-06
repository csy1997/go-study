package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	i := 0
	for {
		<-ticker.C
		i++
		fmt.Println("i =", i)

		if i == 5 {
			ticker.Stop()
			break//不加break会取不出报错
		}
	}
}