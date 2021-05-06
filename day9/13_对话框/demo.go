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
	_, err1 := builder.AddFromFile("day9/13_对话框/test_glade.glade")
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

	//问题对话框
	button1.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win,              //指定父窗口
			gtk.DIALOG_MODAL, //模态对话框（对话框出来前面的窗口没法点）
			gtk.MESSAGE_QUESTION, //问题对话框
			gtk.BUTTONS_YES_NO,   //按钮，可选择是或否
			"这是问题对话框",        //内容
		)
		dialog.SetTitle("问题对话框")

		//运行然后销毁
		ret := dialog.Run()
		//根据点击是还是否，返回不同的值处理
		if ret == gtk.RESPONSE_YES {
			fmt.Println("yes")
		} else if ret == gtk.RESPONSE_NO {
			fmt.Println("no")
		} else {
			fmt.Println("close")
		}
		dialog.Destroy()
	})

	//消息对话框
	button2.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win,              //指定父窗口
			gtk.DIALOG_MODAL, //模态对话框（对话框出来前面的窗口没法点）
			gtk.MESSAGE_INFO, //消息对话框
			gtk.BUTTONS_OK,   //按钮
			"这是消息对话框",        //内容
		)
		dialog.SetTitle("消息对话框")

		//运行然后销毁
		dialog.Run()
		dialog.Destroy()
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
