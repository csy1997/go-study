package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//var start, end int
	//fmt.Println("请输入起始页（>=1）：")
	//fmt.Scan(&start)
	//fmt.Println("请输入终止页：")
	//fmt.Scan(&end)

	//"http://xiaohua.zol.com.cn/"
	DoWork3("http://xiaohua.zol.com.cn/")
}

func DoWork3(url string) {
	//爬
	res, err1 := HttpGet3(url)
	if err1 != nil {
		fmt.Println("httpGet err =", err1)
		return
	}
	//取
	//解释表达式，找`href="/detail(...).html"`
	re := regexp.MustCompile(`href="/detail(?s:(.*?)).html"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	//取关键信息
	jokeUrls := re.FindAllStringSubmatch(res, -1)
	//fmt.Println("jokeUrls =", jokeUrls)
	//fmt.Println("len =", len(jokeUrls))

	//fileTitle := make([]string, 0)
	//fileContent := make([]string, 0)

	page := make(chan int)

	//取网址
	for i, data := range jokeUrls {
		data := data
		i := i
		go func() {//并发爬和写
			jokeUrl := "http://xiaohua.zol.com.cn/detail" + data[1] + ".html"
			//开始爬笑话
			title, content, err2 := SpiderOneJoke(jokeUrl)
			if err2 != nil {
				fmt.Println("", err2)
				return
			}
			//fmt.Printf("title = #%v#\n", title)
			//fmt.Printf("content = #%v#\n", content)
			//fileTitle = append(fileTitle, title)
			//fileContent = append(fileContent, content)

			//把内容存到文件
			StoreJokeToFile(i, title, content)

			page <- i
		}()
	}

	//把内容存到文件
	//for i := 0; i < len(fileTitle); i++ {
	//	StoreJokeToFile(i, fileTitle, fileContent)
	//}

	for i := 0; i < len(jokeUrls); i++ {
		fmt.Printf("第%d个页面爬取并写入完成\n", <-page)
	}
}

func StoreJokeToFile(i int, title string, content string) {
	filename := "day8/08_段子爬虫/" + strconv.Itoa(i+1) + ".txt"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("err1 =", err1)
		return
	}
	defer f.Close()

	//写标题和内容
	f.WriteString(title + "\n\n")
	f.WriteString(content)
}

func SpiderOneJoke(url string) (title, content string, err error){
	//爬
	res, err1 := HttpGet3(url)
	if err1 != nil {
		err = err1
		fmt.Println("httpGet err =", err)
		return
	}

	//取关键信息
	//取标题，只取一个，匹配`name="keywords" content="..."`
	re1 := regexp.MustCompile(`name="keywords" content="(?s:(.*?))"`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile error")
		return
	}
	//取内容
	tmpTitles := re1.FindAllStringSubmatch(res, 1)
	for _, data := range tmpTitles {
		title = data[1]
		break//只取一个
	}

	//取正文，只取一个，匹配`name="description" content="..."`
	re2 := regexp.MustCompile(`name="description" content="(?s:(.*?))"`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile error")
		return
	}
	//取内容
	tmpContents := re2.FindAllStringSubmatch(res, 1)
	for _, data := range tmpContents {
		content = data[1]
		break//只取一个
	}

	return
}

func HttpGet3(url string) (res string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("err = ", err)
		return
	}
	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("err2 =", err2)
			break
		}
		res += string(buf[:n])
	}
	res = ConvertToString(res, "gbk", "utf-8")
	return
}

//gbk转utf-8
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}
