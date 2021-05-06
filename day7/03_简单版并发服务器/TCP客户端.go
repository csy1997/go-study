package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer conn.Close()

	//接收服务器回复的数据
	go func() {
		//从键盘输入内容，给服务器发送消息
		buf := make([]byte, 1024)
		for {
			n, err2 := os.Stdin.Read(buf)
			if err2 != nil {
				fmt.Println("os.Stdin err =", err2)
				return
			}

			//发送给服务器
			conn.Write(buf[:n])
		}

	}()

	//切片缓存，接收服务器回复的消息
	buf2 := make([]byte, 1024)
	for {
		n, err2 := conn.Read(buf2)
		if err2 != nil {
			fmt.Println("conn.Read err =", err2)
			return//conn断开，主协程return，其他携程（读打字内容发送）也终止，程序结束
		}
		fmt.Println("收到服务器的消息：", string(buf2[:n]))
	}
}
