package main

import (
	"encoding/json"
	"fmt"
)

type IT2 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `
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
	`
	var tmp1 IT2
	err1 := json.Unmarshal([]byte(jsonBuf), &tmp1) //第二个参数要地址传递
	if err1 != nil {
		fmt.Println("err1 =", err1)
	} else {
		fmt.Printf("tmp1 = %+v\n", tmp1)
	}

	type IT3 struct {
		Subjects []string `json:"subjects"`
	}
	//可以映射到只包含Subjects的IT3上
	var tmp2 IT3
	err2 := json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数要地址传递
	if err2 != nil {
		fmt.Println("err2 =", err2)
	} else {
		fmt.Printf("tmp2 = %+v\n", tmp2)
	}
}
