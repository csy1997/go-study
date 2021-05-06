package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//不需要写数据时，关闭channel
		close(ch)
		//ch <- 666//关闭channel后无法再发数据，否则会报panic（不过可以读）
	}()

	for {
		//如果ok为true，说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num =", num)
		} else {
			break
		}
	}

	//直接用for range也可以关闭时自动跳出循环
	//for num := range ch {
	//	fmt.Println("num =", num)
	//}
}
