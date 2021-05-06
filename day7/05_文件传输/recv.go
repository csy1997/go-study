package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("net.Listen err =", err1)
		return
	}

	defer listener.Close()

	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("listener.Accept err", err2)
		return
	}

	defer conn.Close()

	//读取对方发送的文件名
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err =", err3)
		return
	}
	fileName := string(buf[:n])
	fmt.Println("fileName =", fileName)

	//回复"ok"
	conn.Write([]byte("ok"))

	var newPath string
	fmt.Scan(&newPath)
	//接收文件内容
	RecvFile(newPath, conn)
}

func RecvFile(path string, conn net.Conn) {
	//新建文件
	f, err1 := os.Create(path)
	if err1 != nil {
		fmt.Println("os.Create err =", err1)
		return
	}
	//defer f.Close()

	//读多少写多少
	buf := make([]byte, 1024)
	for {
		//接收对方发过来的文件内容
		n, err2 := conn.Read(buf)
		if err2 != nil {
			if err2 == io.EOF {
				fmt.Println("文件全部接收完")
			} else {
				fmt.Println("conn.Read =", err2)
			}
			return
		}

		//往文件写入内容
		_, err3 := f.Write(buf[:n])
		if err3 != nil {
			fmt.Println("f.Write err =", err3)
		}
	}
}
