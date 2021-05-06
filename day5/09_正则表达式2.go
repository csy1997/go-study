package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 agsdg 1.23 7. 8.99 lsdljgl 6.66"
	//1.解释规则，小数
	//reg1 := regexp.MustCompile(`\d\.\d`)
	reg1 := regexp.MustCompile(`\d+\.\d+`)//"+"匹配前一个字符的一次或多次
	if reg1 == nil {
		fmt.Println("err =", reg1)
		return
	}
	//2.根据规则提取关键信息
	res1 := reg1.FindAllStringSubmatch(buf, -1)//找所有满足的
	fmt.Println("res1 =", res1)
}
