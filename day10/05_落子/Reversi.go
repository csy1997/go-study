package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

//控件结构体
type ChessWidget struct {
	window      *gtk.Window //窗口
	buttonMin   *gtk.Button //最小化按钮
	buttonClose *gtk.Button //关闭按钮
	labelBlack  *gtk.Label  //黑棋数
	labelWhite  *gtk.Label  //白棋数
	labelTime   *gtk.Label  //倒计时
	imageBlack  *gtk.Image  //黑棋落子
	imageWhite  *gtk.Image  //白棋落子
}

//控件属性结构体
type ChessInfo struct {
	w, h           int //窗口宽度和高度
	x, y           int //鼠标点击的坐标
	startX, startY int //棋盘在窗口的起点坐标（左上）
	gridW, gridH   int //棋盘每格的宽和高
}

//黑白棋结构体
type ChessBoard struct {
	ChessWidget
	ChessInfo

	currentRole int //该谁落子
	tipTimerId  int //实现提示闪烁效果的定时器id

	chess [8][8]int //二维数组，标志棋盘状态
}

//枚举，标志黑白棋子状态
const (
	Empty = iota //当前棋盘格子没有子
	Black        //当前格子为黑子
	White        //当期格子为白子
)

//给按钮设置图标
func ButtonSetImageFromFile(button *gtk.Button, filename string) {
	//获取按钮的大小
	w, h := 0, 0
	button.GetSizeRequest(&w, &h)

	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)

	//创建image
	image := gtk.NewImageFromPixbuf(pixbuf)

	//释放pixbuf资源
	pixbuf.Unref()

	//button设置image
	button.SetImage(image)

	//去掉按钮的焦距
	button.SetCanFocus(false)
}

//给image设置图片
func ImageSetPicFromFile(image *gtk.Image, filename string) {
	//获取按钮的大小
	w, h := 0, 0
	image.GetSizeRequest(&w, &h)

	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)

	//给image设置图片
	image.SetFromPixbuf(pixbuf)

	//释放pixbuf资源
	pixbuf.Unref()
}

//创建控件，设置控件属性
func (obj *ChessBoard) CreateWindow() {
	//加载glade文件
	builder := gtk.NewBuilder()
	_, err1 := builder.AddFromFile("day10/05_落子/test_glade.glade")
	if err1 != nil {
		fmt.Println("err1 =", err1)
	}

	//窗口相关
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

	//按钮相关
	//获取最小化和关闭按钮控件，赋给obj成员
	obj.buttonMin = gtk.ButtonFromObject(builder.GetObject("buttonMin"))
	obj.buttonClose = gtk.ButtonFromObject(builder.GetObject("buttonClose"))

	//给按钮设置图片
	ButtonSetImageFromFile(obj.buttonMin, "day10/image/最小化.png")
	ButtonSetImageFromFile(obj.buttonClose, "day10/image/关闭.png")

	//标签相关
	obj.labelBlack = gtk.LabelFromObject(builder.GetObject("labelBlack"))
	obj.labelWhite = gtk.LabelFromObject(builder.GetObject("labelWhite"))
	obj.labelTime = gtk.LabelFromObject(builder.GetObject("labelTime"))

	//设置字体大小
	obj.labelBlack.ModifyFontSize(30)
	obj.labelWhite.ModifyFontSize(30)
	obj.labelTime.ModifyFontSize(30)

	//设置内容
	obj.labelBlack.SetText("2")
	obj.labelWhite.SetText("2")
	obj.labelTime.SetText("30")

	//改变字体颜色
	obj.labelBlack.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("red"))
	obj.labelWhite.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("red"))
	obj.labelTime.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("black"))

	//image相关
	obj.imageBlack = gtk.ImageFromObject(builder.GetObject("imageBlack"))
	obj.imageWhite = gtk.ImageFromObject(builder.GetObject("imageWhite"))

	//设置图片
	ImageSetPicFromFile(obj.imageBlack, "day10/image/黑棋子.png")
	ImageSetPicFromFile(obj.imageWhite, "day10/image/白棋子.png")

	//棋盘坐标相关
	obj.startX, obj.startY = 200, 100
	obj.gridW, obj.gridH = 50, 40
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

	//判断鼠标按下位置是否在棋盘上，具体在哪个格子
	i := mod(obj.x-obj.startX, obj.gridW)
	j := mod(obj.y-obj.startY, obj.gridH)
	//i := (obj.x - obj.startX) / obj.gridW
	//j := (obj.y - obj.startY) / obj.gridH

	if i < 0 || i > 7 || j < 0 || j > 7 {
		fmt.Println("点到棋盘外")
	} else {
		fmt.Printf("点到格子坐标：(%v, %v)\n", i, j)
		obj.chess[i][j] = Black
		//放棋子后刷图
		obj.window.QueueDraw()
	}
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

//绘图事件的函数
func PaintEvent(ctx *glib.CallbackContext) {
	data := ctx.Data()
	obj, ok := data.(*ChessBoard)
	if ok == false {
		fmt.Println("MouseMoveEvent ChessBoard err")
		return
	}

	//获取画家，设置绘图区域
	painter := obj.window.GetWindow().GetDrawable()
	gc := gdk.NewGC(painter)

	//新建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("day10/image/背景.jpg", obj.w, obj.h, false)

	//黑白棋pixbuf
	chessBlackPixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("day10/image/黑棋子.png", obj.gridW, obj.gridH, false)
	chessWhitePixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("day10/image/白棋子.png", obj.gridW, obj.gridH, false)

	//画背景图
	painter.DrawPixbuf(gc, pixbuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	//画棋子
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if obj.chess[i][j] == Black {
				painter.DrawPixbuf(gc, chessBlackPixbuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			} else if obj.chess[i][j] == White {
				painter.DrawPixbuf(gc, chessWhitePixbuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			}
		}
	}

	//释放资源
	pixbuf.Unref()
	chessBlackPixbuf.Unref()
	chessWhitePixbuf.Unref()
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

	//按钮的信号处理
	obj.buttonClose.Clicked(func() {
		//关闭窗口前先关闭定时器
		glib.TimeoutRemove(obj.tipTimerId)

		gtk.MainQuit() //关闭窗口
	})

	obj.buttonMin.Clicked(func() {
		obj.window.Iconify() //最小化窗口
	})

	//绘图相关
	//大小改变事件
	obj.window.Connect("configure_event", func() {
		//重新刷图
		obj.window.QueueDraw()
	})

	//绘图事件
	obj.window.Connect("expose-event", PaintEvent, obj)
}

//黑白棋属性相关
func (obj *ChessBoard) InitChess() {
	//初始都先隐藏
	obj.imageBlack.Hide()
	obj.imageWhite.Hide()

	//黑子先下，初始落子设为黑子
	obj.currentRole = Black

	//启动定时器
	obj.tipTimerId = glib.TimeoutAdd(500, func() bool {
		showTip(obj)
		return true
	})
}

//提示功能，实现闪烁效果
func showTip(obj *ChessBoard) {
	if obj.currentRole == Black { //当前黑子下
		//隐藏白子image
		obj.imageWhite.Hide()
		//实现闪烁
		if obj.imageBlack.GetVisible() { //原来是显示的，现在要隐藏
			obj.imageBlack.Hide()
		} else { //原来是隐藏的，现在要显示
			obj.imageBlack.Show()
		}
	} else { //当前白子下
		//隐藏黑子image
		obj.imageBlack.Hide()
		//实现闪烁
		if obj.imageWhite.GetVisible() { //原来是显示的，现在要隐藏
			obj.imageWhite.Hide()
		} else { //原来是隐藏的，现在要显示
			obj.imageWhite.Show()
		}
	}
}

func mod(x, y int) (res int) {
	if x < 0 || y <= 0 {
		res = -1
	} else {
		res = x / y
	}
	return
}

func main() {
	//初始化
	gtk.Init(&os.Args)

	var obj ChessBoard

	obj.CreateWindow() //创建控件，设置控件属性
	obj.HandleSignal() //事件、信号处理

	obj.InitChess() //黑白棋相关初始化

	obj.window.Show() //显示控件

	//主事件循环
	gtk.Main()
}
