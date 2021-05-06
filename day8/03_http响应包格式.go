package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err1 := net.Dial("tcp", ":8080")
	if err1 != nil {
		fmt.Println("err1 =", err1)
		return
	}
	defer conn.Close()

	requestBuf := "GET / HTTP/1.1\nHost: 127.0.0.1:8080\nConnection: keep-alive\nUpgrade-Insecure-Requests: 1\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\nSec-Fetch-Site: none\nSec-Fetch-Mode: navigate\nSec-Fetch-User: ?1\nSec-Fetch-Dest: document\nAccept-Encoding: gzip, deflate, br\nAccept-Language: zh-CN,zh;q=0.9"

	//先发请求包，服务器才会回响应包
	conn.Write([]byte(requestBuf))

	//接收服务器回复的响应包
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("Read err =", err2)
		return
	}

	//打印相应报文
	fmt.Printf("#%v#", string(buf[:n]))
}