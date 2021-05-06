package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件")
	var path string
	fmt.Scan(&path)

	//获取文件名
	info, err1 := os.Stat(path)
	if err1 != nil {
		fmt.Println("os.Stat err1 =", err1)
		return
	}
	//主动连接服务器
	conn, err2 := net.Dial("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println("net.Dial err", err2)
		return
	}

	defer conn.Close()

	//给接收方，先发送文件名
	_, err3 := conn.Write([]byte(info.Name()))
	if err3 != nil {
		fmt.Println("conn.Write err", err3)
		return
	}

	//接收对方的回复，如果对方回复"ok"，说明对方准备好，可以发送文件了
	var n int
	buf := make([]byte, 1024)
	n, err4 := conn.Read(buf)
	if err4 != nil {
		fmt.Println("conn.Read err", err4)
		return
	}
	if "ok" == string(buf[:n]) {
		//发送文件内容
		SendFile(path, conn)
	}
}

func SendFile(path string, conn net.Conn) {
	//defer conn.Close()
	//以只读方式打开文件
	f, err1 := os.Open(path)
	if err1 != nil {
		fmt.Println("os.Open err", err1)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)

	//读文件内容，读多少发多少
	for {
		//从文件读
		n, err2 := f.Read(buf)
		fmt.Println("n ==", n)
		if err2 != nil {
			if err2 == io.EOF {//文件读完
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err", err2)
			}
			return
		}
		//写进conn，发给服务器
		conn.Write(buf[:n])
	}
}
