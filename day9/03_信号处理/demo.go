package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
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
	button2.SetSizeRequest(100, 100) //设置button2大小
	//6.按钮添加到布局中
	layout.Put(button1, 0, 0)
	layout.Put(button2, 100, 100)
	//7.显示控件
	//如果有多个控件，用show的话，每个控件都要加show
	//win.Show()
	//layout.Show()
	//button1.Show()
	win.ShowAll() //显示目标内所有控件

	//8.信号处理

	//"clicked"代表按下按钮信号
	str := "are you ok?"
	//告诉系统，只要按下按钮就会执行HandleSignal函数，str是该函数参数
	//Connect只会调用一次，之后系统就知道规则了
	button1.Connect("clicked", HandleSignal, str)

	//处理函数可以是匿名函数，推荐写法
	//button2.Connect("clicked", func() {
	//	fmt.Println("---------------------")
	//	fmt.Println("按钮2被按下：", str)
	//	fmt.Println("---------------------")
	//})
	button2.Clicked(func() {
		fmt.Println("---------------------")
		fmt.Println("按钮2被按下：", str)
		fmt.Println("---------------------")
	})

	//窗口的关闭按钮，触发"destroy"
	win.Connect("destroy", func() {
		//fmt.Println("233333333333")
		gtk.MainQuit()//gtk程序关闭
	})

	//主事件循环
	gtk.Main()
	fmt.Println("程序结束")
}

//func HandleSignal(tmp string) {
//	fmt.Println("tmp =", tmp)
//}

func HandleSignal(ctx *glib.CallbackContext) {
	fmt.Println("---------------------")
	arg := ctx.Data() //获取用户传的参数，空接口类型
	data, ok := arg.(string)
	if ok {
		fmt.Println("按钮1被按下：", data)
	}
	fmt.Println("---------------------")
}
