package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "./demo.txt"
	WriteFile(path)
	ReadFile(path)
	ReadFileLine(path)
}

//每次只读取一行
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用完需关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放到缓冲区
	r := bufio.NewReader(f)
	for {
		//遇到\n读取结束，会把\n也读进去
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {//文件已读完
				break
			}
			fmt.Println("err =", err)
		}
		fmt.Printf("buf = \"%s\"\n", string(buf))
	}
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用完需关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {//文件出错，且没有到结尾
		fmt.Println("err =", err)
	}
	fmt.Printf("buf = \n%s", string(buf[:n]))
}

func WriteFile(path string) {
	//打开文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用完需关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		//把"i = i\n"，这个字符串存到buf中
		buf = fmt.Sprintf("i = %d\n", i)

		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err =", err)
		}
		fmt.Println("n =", n)
	}
}
