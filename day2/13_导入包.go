package main

import . "fmt"//加上"."调用函数无需前面写该包名
import io "fmt"//起别名，给fmt起io的名字
import _ "os"//忽略这个包，但会执行该包的init函数

func main() {
	Println("This is a 23_同目录调用")
	io.Println("This is a 23_同目录调用")
}
