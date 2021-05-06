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

	//加载glade文件
	builder := gtk.NewBuilder()
	_, err1 := builder.AddFromFile("day9/04_glade的使用/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//获取glade文件里的控件
	win := gtk.WindowFromObject(builder.GetObject("window1")) //根据glade文件里对应控件的名称来获取
	win.SetTitle("嘻嘻嘻")
	win.SetIconFromFile("day9/05_窗口/test_jpg.jpg")//设置图标
	win.SetSizeRequest(480, 320)//设置窗口最小大小
	//获取窗口大小
	var w, h int
	win.GetSize(&w, &h)
	fmt.Printf("w = %d, h = %d\n", w, h)
	//让窗口居中显示
	win.SetPosition(gtk.WIN_POS_CENTER)
	//固定窗口大小
	win.SetResizable(false)

	//信号处理
	//button1.Clicked(func() {
	//	fmt.Println("按钮111被按下")
	//})
	//button2.Clicked(func() {
	//	fmt.Println("按钮222被按下")
	//})
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//显示控件，如果是通过glade添加的控件，show即可显示所有（如果是代码布局必须用showAll）
	win.Show()

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
