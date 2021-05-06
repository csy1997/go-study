package main

import (
	"fmt"
	"net/http"
)

func main() {
	//resp, err1 := http.Get("http://www.baidu.com")
	resp, err1 := http.Get("http://127.0.0.1:8080")
	if err1 != nil {
		fmt.Println("err1 =", err1)
		return
	}
	defer resp.Body.Close()

	fmt.Println("status =", resp.Status)
	fmt.Println("statusCode =", resp.StatusCode)
	fmt.Println("header =", resp.Header)
	//fmt.Println("body =", resp.Body)//body需要读出来

	buf := make([]byte, 1024)
	var tmp string
	for {
		n, err2 := resp.Body.Read(buf)
		if err2 != nil {
			fmt.Println("read err =", err2)
			break
		}
		tmp += string(buf[:n])
	}

	fmt.Println("tmp =", tmp)
}
