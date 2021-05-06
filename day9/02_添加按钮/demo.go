package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//1.创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	//2.设置窗口属性
	win.SetTitle("呵呵呵")
	win.SetSizeRequest(480, 320)
	//3.创建容器控件（固定布局，任意布局）
	layout := gtk.NewFixed()
	//4.布局添加到窗口上
	win.Add(layout)
	//5.创建按钮
	button1 := gtk.NewButtonWithLabel("^_^")
	button2 := gtk.NewButtonWithLabel("@_@")
	button2.SetSizeRequest(100, 100)//设置button2大小
	//6.按钮添加到布局中
	layout.Put(button1, 0, 0)
	layout.Put(button2, 100, 100)
	//7.显示控件
	//如果有多个控件，用show的话，每个控件都要加show
	//win.Show()
	//layout.Show()
	//button1.Show()
	win.ShowAll()//显示目标内所有控件

	//主事件循环
	gtk.Main()
}
