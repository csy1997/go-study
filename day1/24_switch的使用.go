package main

import "fmt"

func main() {
	num := 1

	switch num {
	case 1:
		fmt.Println("按下的是1楼")
		//break//默认包含了break
		fallthrough //加上不跳出switch语句，后面的无条件进入执行
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("按下的楼不存在")
	}

	//可以放一个初始化语句
	switch s := "e"; s {
	case "a":
		fmt.Println("是a")
	case "b":
		fmt.Println("是b")
	case "c", "d", "e"://可以多个选项满足一个即可
		fmt.Println("是c或d或e")
	}

	score := 50
	switch {//可以不加条件
	case score > 50://case里放条件
		fmt.Println("优秀")
	case score > 40:
		fmt.Println("良好")
	case score > 30:
		fmt.Println("及格")
	}
}
