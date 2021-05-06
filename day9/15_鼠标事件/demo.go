package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

func main() {
	//初始化
	gtk.Init(&os.Args)

	//加载glade文件
	builder := gtk.NewBuilder()
	_, err1 := builder.AddFromFile("day9/15_鼠标事件/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//获取glade文件里的控件
	win := gtk.WindowFromObject(builder.GetObject("window1")) //根据glade文件里对应控件的名称来获取

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//--------------------------------------------------

	//添加鼠标按下事件
	//BUTTON_PRESS_MASK：鼠标按下，触发信号"button-press-event"
	//BUTTON_MOTION_MASK：鼠标移动，按下任何键都可以
	win.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))//Button1代表左键

	//"button-press-event" 鼠标按下时触发
	win.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		//获取按下鼠标键属性结构体变量，是系统内部的变量，不是用户传的参
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		//fmt.Println("按下鼠标")
		//左中右键分别对应1,2,3
		fmt.Println("按下鼠标键", event.Button)

		if event.Type == int(gdk.BUTTON_PRESS) {
			fmt.Println("单击")
		} else if event.Type == int(gdk.BUTTON2_PRESS) {
			fmt.Println("双击")
		}

		fmt.Printf("相对于窗口坐标：%v, %v\n", event.X, event.Y)
		fmt.Printf("相对于屏幕坐标：%v, %v\n", event.XRoot, event.YRoot)
	})

	//"motion-notify-event" 按住鼠标移动时触发
	win.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		fmt.Printf("按住鼠标移动-----------相对于窗口坐标：%v, %v\n", event.X, event.Y)
		fmt.Printf("按住鼠标移动-----------相对于屏幕坐标：%v, %v\n", event.XRoot, event.YRoot)
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
