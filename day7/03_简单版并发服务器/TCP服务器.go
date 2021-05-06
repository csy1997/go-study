package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer listener.Close()

	//阻塞等待用户链接
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("err1 =", err)
			continue
		}
		//处理用户请求
		go HandleConn(conn)//每来一个处理一个，且不影响后面，需创建协程
	}
}

//处理用户请求
func HandleConn(conn net.Conn) {
	//调用完自动关闭conn
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("连接成功", addr)

	buf := make([]byte, 2048)

	//循环读取连接用户数据
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err =", err)
			return
		}
		msg := string(buf[:n])
		fmt.Printf("读取%s的长为%d数据：%s\n", addr, len(msg), msg)

		if "exit" == string(buf[:n-1]) {//需要减掉一个换行符
			fmt.Println(addr, "已断开")
			return
		}

		//把数据转换为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(msg)))
	}
}
