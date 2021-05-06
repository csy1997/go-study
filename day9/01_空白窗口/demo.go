package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化（固定）
	gtk.Init(&os.Args)

	//用户写的代码
	//1.创建窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	//2.设置属性（标题，大小）
	win.SetTitle("go gtk")
	win.SetSizeRequest(480, 320)
	//3.显示窗口
	win.Show()

	//主事件循环（固定）
	//让程序不结束，等待用户操作
	//fmt.Println("!!!")
	gtk.Main()
	//fmt.Println("...")
}
