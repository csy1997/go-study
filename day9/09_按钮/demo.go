package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	_, err1 := builder.AddFromFile("day9/09_按钮/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//获取glade文件里的控件
	win := gtk.WindowFromObject(builder.GetObject("window1")) //根据glade文件里对应控件的名称来获取

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//获取按钮控件
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	button2 := gtk.ButtonFromObject(builder.GetObject("button2"))

	button1.SetLabel("^_^")//设置文本内容
	button1.SetLabelFontSize(30)//改变字体大小
	fmt.Println("button1 text =", button1.GetLabel())
	button1.SetSensitive(false)//使变灰，不让按

	//获取button2大小，输出到w,h
	var w, h int
	button2.GetSizeRequest(&w, &h)

	//新建图片资源
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("day9/09_按钮/testjpg.jpg", w-10, h-10, false)

	//新建image
	image := gtk.NewImageFromPixbuf(pixbuf)
	pixbuf.Unref()//释放资源

	button2.SetImage(image)

	button2.SetCanFocus(false)//取消焦距（外面的框）

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
