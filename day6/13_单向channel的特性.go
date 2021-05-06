package main

func main() {
	//创建一个双向的channel
	ch := make(chan int)

	//双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch//只能写不能读
	var readCh <-chan int = ch//只能读不能写

	writeCh <- 666
	//<-writeCh
	<-readCh
	//readCh <- 666

	//单向无法转为双向
	//var ch2 chan int = writeCh
}
