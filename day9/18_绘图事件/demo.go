package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
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
	_, err1 := builder.AddFromFile("day9/18_绘图事件/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//获取glade文件里的控件
	win := gtk.WindowFromObject(builder.GetObject("window1")) //根据glade文件里对应控件的名称来获取

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//--------------------------------------------------

	var w, h int
	//"configure_event" 键盘按键按下时触发
	win.Connect("configure_event", func() {
		win.GetSize(&w, &h)

		//改变窗口大小时绘的图不会变，因此会出现重叠
		//可以让窗口改变时刷新绘图，即触发"expose-event"，图片可以随窗口大小自适应
		win.QueueDraw()
	})

	//允许窗口绘图
	win.SetAppPaintable(true)

	x := 0

	//"expose-event" 绘图时触发信号
	win.Connect("expose-event", func() {
		//设置画家，指定绘图区域
		painter := win.GetWindow().GetDrawable()
		gc := gdk.NewGC(painter)

		//创建图片资源
		bg, _ := gdkpixbuf.NewPixbufFromFileAtScale("day9/18_绘图事件/testjpg.jpg", w, h, false)
		face, _ := gdkpixbuf.NewPixbufFromFileAtScale("day9/18_绘图事件/myy.jpg", 80, 80, false)

		//开始绘图（覆盖全窗口）
		painter.DrawPixbuf(gc, bg, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		//脸图横向移动随x变化，纵向不变
		painter.DrawPixbuf(gc, face, 0, 0, x, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

		//释放图片资源
		bg.Unref()
		face.Unref()
	})

	//获取按钮
	button := gtk.ButtonFromObject(builder.GetObject("button1"))
	//按按钮实现脸图在窗口中移动
	button.Clicked(func() {
		x += 50
		//脸图到边缘时回到最左，循环
		if x >= w {
			x =0
		}

		//刷图
		win.QueueDraw()
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
