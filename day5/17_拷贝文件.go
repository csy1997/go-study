package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args//获取命令行参数

	if len(list) != 3 {
		fmt.Println("usage: 路径 srcFile dstFile")
		return
	}
	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}
	//只读方式打开源文件
	srcFile, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 =", err1)
		return
	}

	//新建目的文件
	dstFile, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 =", err2)
		return
	}

	//操作完文件，需关闭
	defer srcFile.Close()
	defer dstFile.Close()

	//拷贝，读多少写多少
	buf := make([]byte, 4*1024)
	for {
		n, err := srcFile.Read(buf)//从源文件中读
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件已读完")
				break
			}
			fmt.Println("err =", err)
		}
		dstFile.Write(buf[:n])
	}
}
