package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)

	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10)//10代表转为10进制类型再追加
	slice = strconv.AppendQuote(slice, "go")

	fmt.Println(slice)
	fmt.Println(string(slice))

	//其他类型转为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println(str)
	//'f'指打印格式（小数方式），-1指小数点位数（紧缩模式），64指以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(str)

	//整形转字符串，常用
	str = strconv.Itoa(6666)
	fmt.Sprintln(str)

	//字符串转其他类型
	flag, err := strconv.ParseBool("true")
	//flag, err := strconv.ParseBool("tru")
	if err == nil {
		fmt.Println(flag)
	} else {
		fmt.Println(err)
	}

	//字符串转整形，常用
	a, err1 := strconv.Atoi("1234")
	//a, err1 := strconv.Atoi("1234a")
	if err1 == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err1)
	}
}
