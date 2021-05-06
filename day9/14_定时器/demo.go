package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	_, err1 := builder.AddFromFile("day9/14_定时器/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//获取glade文件里的控件
	win := gtk.WindowFromObject(builder.GetObject("window1")) //根据glade文件里对应控件的名称来获取

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//--------------------------------------------------

	//获取按钮
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	button2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	button2.SetSensitive(false)//初始让停止按钮变灰

	//获取标签
	label := gtk.LabelFromObject(builder.GetObject("label1"))
	label.SetText("0")
	label.ModifyFontSize(30)

	num := 0
	id := 0

	//启动定时器
	button1.Clicked(func() {
		id = glib.TimeoutAdd(1000, func() bool {
			//每隔一秒显示一个数
			num++
			label.SetText(strconv.Itoa(num))//转成字符串
			return true//必须要return true，才能一直循环，而不是只执行一次
		})

		button1.SetSensitive(false)//开始计时启动键变灰
		button2.SetSensitive(true)
	})

	//停止计时器
	button2.Clicked(func() {
		glib.TimeoutRemove(id)

		button1.SetSensitive(true)
		button2.SetSensitive(false)//停止计时停止键变灰
	})

	//通过代码添加控件，需要显示所有
	win.ShowAll()

	//--------------------------------------------

	//显示控件，如果是通过glade添加的控件，show即可显示所有（如果是代码布局必须用showAll）
	//win.Show()

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
