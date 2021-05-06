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

	DoWork2(start, end)

}

func DoWork2(start int, end int) {
	fmt.Printf("正在爬取%d到%d的页面", start, end)

	page := make(chan int)

	//https://tieba.baidu.com/p/5545487283?pn=1
	for i := start; i <= end; i++ {
		go SpiderPape(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬完\n", <-page)//谁先爬完就会把i放入管道，就会打印，数量总共和上面一样
	}
}

func SpiderPape(i int, page chan int)  {
	url := "https://tieba.baidu.com/p/5545487283?pn=" + strconv.Itoa(i)
	//爬
	res, err1 := HttpGet2(url)
	if err1 != nil {
		fmt.Println("httpGet err =", err1)
		return
	}
	//把内容写入文件
	//dir, err2 := os.Getwd()
	//if err2 != nil {
	//	fmt.Println("os.Getwd err =", err2)
	//	return
	//}
	os.MkdirAll("day8/07_并发版爬虫", os.ModePerm)
	filename := "day8/07_并发版爬虫/" + strconv.Itoa(i) + ".html"
	f, err2 := os.Create(filename)
	if err2 != nil {
		fmt.Println("os.Create err =", err2)
		return
	}
	f.WriteString(res)

	f.Close()

	page <- i
}

//爬取网页内容
func HttpGet2(url string) (res string, err error){
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