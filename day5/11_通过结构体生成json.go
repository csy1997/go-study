package main

import (
	"encoding/json"
	"fmt"
)

/*
{
	"company": "itcast",
	"subjects": [
		"Go",
		"C++",
		"Python",
		"Test"
	],
	"isok": true,
	"price": 666.666
}
*/

type IT struct {
	//成员变量首字母必须大写
	Company  string   `json:"company"` //二次编码，可让json输出格式变为小写
	Subjects []string `json:"-"`       //可使字段不输出到屏幕
	IsOk     bool     `json:",string"` //使其以string格式输出
	Price    float64  `json:",string"`
}

func main() {
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	//编码，根据内容生成json文本
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("buf =", string(buf))
	}

}
