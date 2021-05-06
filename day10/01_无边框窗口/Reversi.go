package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

//控件结构体
type ChessWidget struct {
	window *gtk.Window
}

//控件属性结构体
type ChessInfo struct {
	w, h int
	x, y int
}

//黑白棋结构体
type ChessBoard struct {
	ChessWidget
	ChessInfo
}

//创建控件，设置控件属性
func (obj *ChessBoard) CreateWindow() {
	//加载glade文件
	builder := gtk.NewBuilder()
	_, err1 := builder.AddFromFile("day10/01_无边框窗口/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//获取目标窗口，赋给成员window
	obj.window = gtk.WindowFromObject(builder.GetObject("window1"))
	//允许绘图
	obj.window.SetAppPaintable(true)
	//设置初始居中显示
	obj.window.SetPosition(gtk.WIN_POS_CENTER)
	//设置窗口的宽和高
	obj.w, obj.h = 800, 480
	obj.window.SetSizeRequest(800, 480)
	//去边框
	obj.window.SetDecorated(false)

	//设置事件，让窗口可以捕获鼠标点击和移动
	obj.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
}

//鼠标点击事件的函数
func MousePressEvent(ctx *glib.CallbackContext) {
	//先判断用户传的是不是ChessBoard类型的参数
	//获取用户传的参（鼠标点击事件Connect的第三个参数obj），然后类型断言
	data := ctx.Data()
	obj, ok := data.(*ChessBoard)
	if ok == false {
		fmt.Println("MousePressEvent ChessBoard err")
		return
	}

	//获取按下鼠标键属性结构体变量，是系统内部的变量，不是用户传的参
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	//保存点击的x，y坐标（这里obj是地址传递，所以原obj中x和y会被赋值）
	obj.x, obj.y = int(event.X), int(event.Y)
	//fmt.Printf("点击：x = %v, y = %v\n", obj.x, obj.y)
}

//鼠标按下移动事件的函数
func MouseMoveEvent(ctx *glib.CallbackContext) {
	//先判断用户传的是不是ChessBoard类型的参数
	//获取用户传的参（鼠标点击事件Connect的第三个参数obj），然后类型断言
	data := ctx.Data()
	obj, ok := data.(*ChessBoard)
	if ok == false {
		fmt.Println("MouseMoveEvent ChessBoard err")
		return
	}

	//获取按下鼠标键属性结构体变量，是系统内部的变量，不是用户传的参
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	//让窗体跟随按下鼠标的移动而移动，只需实时计算鼠标相对于按下时位置的位移，然后用在窗体上就可以
	x, y := int(event.XRoot), int(event.YRoot)
	obj.window.Move(x-obj.x, y-obj.y)
	//fmt.Printf("移动：XRoot = %v, YRoot = %v\n", x, y)
	//fmt.Printf("     x = %v, y = %v\n", obj.x, obj.y)
}

//控件的信号处理
func (obj *ChessBoard) HandleSignal() {
	//关闭信号
	obj.window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//鼠标点击事件
	obj.window.Connect("button-press-event", MousePressEvent, obj)

	//鼠标按下移动事件
	obj.window.Connect("motion-notify-event", MouseMoveEvent, obj)
}

func main() {
	//初始化
	gtk.Init(&os.Args)

	var obj ChessBoard

	obj.CreateWindow()//创建控件，设置控件属性
	obj.HandleSignal()//事件、信号处理

	obj.window.Show()//显示控件

	//主事件循环
	gtk.Main()
}
