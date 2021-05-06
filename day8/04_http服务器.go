package main

import (
	"fmt"
	"net/http"
)

func HandleConn(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method =", request.Method)
	fmt.Println("header =", request.Header)
	fmt.Println("url =", request.URL)
	fmt.Println("body =", request.Body)

	writer.Write([]byte("hello world"))//给客户端回复数据
}

func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/", HandleConn)

	//监听绑定
	http.ListenAndServe("127.0.0.1:8080", nil)
}
