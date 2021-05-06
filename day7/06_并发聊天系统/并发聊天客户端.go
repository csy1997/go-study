package main

import (
	"fmt"
	"net"
)

var isDisConn = make(chan bool)

func main() {
	//建立连接
	conn, err1 := net.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("net.Dial err =", err1)
	}

	defer conn.Close()

	go func() {
		//收消息
		buf := make([]byte, 1024)
		for {
			n, err2 := conn.Read(buf)
			if n == 0 {//服务器断开或其他问题
				fmt.Println("conn.Read err =", err2)
				fmt.Println("服务器断开，自动退出")
				isDisConn <- true
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	//发消息，普通聊天消息
	go func() {
		var msg string
		for {
			fmt.Scan(&msg)
			if len(msg) > 4 && msg[len(msg)-4:] == "|cmd" {
				//fmt.Println("msg =", msg)
				tmp := []byte(msg[:len(msg)-4])
				//命令消息byte串末尾加个0
				tmp = append(tmp, 1)
				_, err3 := conn.Write(tmp)
				if err3 != nil {
					fmt.Println("conn.Write err =", err3)
					fmt.Println("连接失败，已自动断开")
					continue
				}
				if msg == "quit|cmd" {
					fmt.Println("退出成功")
					isDisConn <- true
					return
				}
			} else {
				//fmt.Println("msg =", msg)
				tmp := []byte(msg)
				//普通消息byte串末尾加个0
				tmp = append(tmp, 0)
				_, err3 := conn.Write(tmp)
				if err3 != nil {
					fmt.Println("conn.Write err =", err3)
					continue
				}
			}

		}
	}()

	<-isDisConn

	////发指令消息，退出、改名等
	//buf := make([]byte, 1024)
	//for {
	//	n, err4 := os.Stdin.Read(buf)
	//	if err4 != nil {
	//		fmt.Println("os.Read err =", err4)
	//		continue
	//	}
	//	//指令消息byte串末尾加个1
	//	tmp := append(buf[:n], 1)
	//	_, err5 := conn.Write(tmp)
	//	if err5 != nil {
	//		fmt.Println("conn.Write err =", err5)
	//		continue
	//	}
	//}
}
