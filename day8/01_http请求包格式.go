package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("err1 =", err1)
		return
	}
	defer listener.Close()

	//阻塞等待用户连接
	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("err2 =", err2)
		return
	}
	defer conn.Close()

	//接收客户端的数据
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("err3 =", err3)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))
}
