package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var start, end int
	fmt.Println("请输入起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Println("请输入终止页：")
	fmt.Scan(&end)

	DoWork(start, end)

}

func DoWork(start int, end int) {
	fmt.Printf("正在爬取%d到%d的页面", start, end)

	//https://tieba.baidu.com/p/5545487283?pn=1
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/p/5545487283?pn=" + strconv.Itoa(i)

		//爬
		res, err1 := HttpGet(url)
		if err1 != nil {
			fmt.Println("httpGet err =", err1)
			continue
		}
		//把内容写入文件
		os.MkdirAll("day8/06_百度贴吧小爬虫", os.ModePerm)
		filename := "day8/06_百度贴吧小爬虫/" + strconv.Itoa(i) + ".html"
		//filename := strconv.Itoa(i) + ".html"
		f, err2 := os.Create(filename)
		if err2 != nil {
			fmt.Println("os.Create err =", err2)
			continue
		}
		f.WriteString(res)

		f.Close()
	}
}

//爬取网页内容
func HttpGet(url string) (res string, err error){
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("err =", err)
		return
	}

	defer resp.Body.Close()

	//读取网页body内容
	buf := make([]byte, 1024)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {//读取结束，或者出问题
			fmt.Println("resp.Body.Read err =", err2)
			break
		}
		res += string(buf[:n])
	}
	return
}