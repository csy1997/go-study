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
	_, err1 := builder.AddFromFile("day9/07_图片控件/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//获取glade文件里的控件
	win := gtk.WindowFromObject(builder.GetObject("window1")) //根据glade文件里对应控件的名称来获取

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//获取图片控件
	image := gtk.ImageFromObject(builder.GetObject("image1"))

	//获取控件大小，输出到w,h
	var w, h int
	image.GetSizeRequest(&w, &h)

	//设置一张图片资源，pixbuf，控件大小和图片一样（设为false代表不保存图片原来的尺寸）
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("day9/07_图片控件/testjpg.jpg", w-10, h-10, false)
	image.SetFromPixbuf(pixbuf)//给image设置图片

	//图片资源使用完毕，需要释放空间（注意）
	pixbuf.Unref()

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
