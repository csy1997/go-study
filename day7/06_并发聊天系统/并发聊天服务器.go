package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("net.Listen err =", err1)
		return
	}
	defer listener.Close()

	go Manager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("listener.Accept err =", err2)
			continue
		}
		go HandleConn(conn) //处理用户连接
	}
}

type Client struct {
	C    chan string
	Name string
	Addr string
}

type Message struct {
	Client
	Msg string
}

//保存在线用户
var onlineMap map[string]Client

var msgChan = make(chan Message)

//var commands = make(chan string)

//新开一个协程，转发消息，只要有消息来了，遍历map给每个成员都发送此消息
func Manager() {
	onlineMap = make(map[string]Client)

	for {
		message := <-msgChan

		//遍历map，给每个Client(在服务器中对应的channel)发送msg
		for addr, client := range onlineMap {
			if message.Addr != addr {
				client.C <- message.MakeMsg()
			}
		}
	}
}

//处理用户连接
func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端的网络地址
	clientAddr := conn.RemoteAddr().String()

	//把新用户加到map，先默认用户名和用户地址一样
	client := Client{make(chan string), clientAddr, clientAddr}

	onlineMap[clientAddr] = client

	//新开一个协程，专门给当前客户端发送信息
	go func(cli Client, con net.Conn) {
		for msg := range cli.C {
			conn.Write([]byte(msg))
		}
	}(client, conn)

	//处理该用户连接，同时要向所有用户广播该用户上线的信息
	//msgChan <- "用户 {"+clientAddr+", "+client.Name+"} 已上线"
	msgChan <- Message{client, "已上线"}

	//判断对方是否主动退出
	isQuit := make(chan bool)
	//判断对方是否有数据发送
	hasData := make(chan bool)

	//新开一个协程，记录用户发过来的数据或者指令
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err1 := conn.Read(buf)
			if n == 0 { //对方断开或者其他问题
				isQuit <- true
				fmt.Println("conn.Read err =", err1)
				return
			}
			tag := buf[n-1]
			msg := string(buf[:n-1])
			//fmt.Printf("msg = %s, len(msg) = %d\n", msg, len(msg))
			switch tag {
			case 0:
				msgChan <- Message{client, msg} //转发此信息
			case 1:
				if msg == "quit" {
					//delete(onlineMap, clientAddr)
					//msgChan <- Message{client, "已退出"}
					////conn.Write([]byte("退出成功！"))
					//return
					isQuit <- true
				} else if len(msg) > 7 && msg[:7] == "rename|" {
					msgChan <- Message{client, "改名为" + msg[7:]}
					client.Name = msg[7:]
					onlineMap[clientAddr] = client
					conn.Write([]byte("改名为" + msg[7:] + "成功！"))
				} else if msg == "userList" {
					conn.Write([]byte("用户列表："))
					for _, info := range onlineMap {
						conn.Write([]byte(info.Addr + ": " + info.Name))
					}
				}
			default:
				continue
			}
			hasData <- true
		}
	}()

	for {
		//通过select监测channel流动
		select {
		case <-isQuit:
			delete(onlineMap, clientAddr)
			msgChan <- Message{client, "已退出"}
			fmt.Println("map size: ", len(onlineMap))
			//conn.Write([]byte("退出成功！"))
			return
		case <-hasData:

		case <-time.After(60 * time.Second): //60s没有isQuit或者hasData，就超时退出
			delete(onlineMap, clientAddr)
			msgChan <- Message{client, "已超时退出"}
			return
		}
	}

	//for {//为了不让上面给客户端发消息的协程结束，本函数也不能结束
	//
	//}
}

func (p *Message) MakeMsg() string {
	return "{" + p.Addr + ", " + p.Name + "}: " + p.Msg
}
