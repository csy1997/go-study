package main

import "fmt"

type ChessBoard struct {

	currentRole int //该谁落子

	chess [8][8]int //二维数组，标志棋盘状态
}

//枚举，标志黑白棋子状态
const (
	Empty = iota //当前棋盘格子没有子
	Black        //当前格子为黑子
	White        //当期格子为白子
)

func main() {
	var obj ChessBoard

	//初始化棋盘
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			obj.chess[i][j] = Empty
		}
	}

	//中间4个为2黑2白
	obj.chess[3][3] = Black
	obj.chess[4][4] = Black
	obj.chess[3][4] = White
	obj.chess[4][3] = White

	//黑子先下，初始落子设为黑子
	obj.currentRole = Black

	obj.Show()
	obj.PlaceChess(5, 3)
	obj.Show()
}

func (obj *ChessBoard) PlaceChess(i int, j int) {
	//八个方向分别看能不能吃子
	tmp := [][]int{{0,1,0,-1,0},{1,1,-1,-1,1}}
	//如果能吃到子将flag置为true
	flag := false

	for _, tm := range tmp {
		for k := 0; k < len(tm)-1; k++ {
			x, y := tm[k], tm[k+1]
			i1, j1 := i, j
			for {
				i1 += x
				j1 += y

				//往当前方向遍历chess，遇到对手颜色棋子则继续遍历
				//到达棋盘外或Empty则终止跳出循环
				if i1 < 0 || i1 > 7 || j1 < 0 || j1 > 7 {
					break
				}

				cur := obj.chess[i1][j1]
				if cur == Empty {
					break
				}

				//遇到自己颜色，则往回把中间的棋子（对手颜色）都变成自己颜色（吃子），然后跳出循环
				if cur == obj.currentRole {
					i1 -= x
					j1 -= y
					for i1 != i || j1 != j {//不包括起始的i,j，需最后再判断
						obj.chess[i1][j1] = obj.currentRole
						//如果有改变（吃到子），则flag置为true
						flag = true
						i1 -= x
						j1 -= y
					}
					break
				}
			}
		}
	}
	if flag {
		obj.chess[i][j] = obj.currentRole
		fmt.Println("已落子")
	} else {
		fmt.Println("当前位置不能吃到子，所以不能落子")
	}
}

func (obj *ChessBoard) Show() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(obj.chess[i][j], " ")
		}
		fmt.Println()
	}
}
