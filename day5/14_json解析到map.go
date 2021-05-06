package main

import (
	"encoding/json"
	"fmt"
)

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

	m := make(map[string]interface{}, 4)
	err1 := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数要地址传递
	if err1 != nil {
		fmt.Println("err1 =", err1)
	} else {
		fmt.Printf("tmp1 = %+v\n", m)
	}

	//取其中的一部分，比如company
	var str string
	//str = string(m["company"])//interface{}不能强转成string，需先进行类型断言
	//fmt.Println("str =", str)

	//类型断言
	for key, value := range m {
		switch data := value.(type) {
		case string:
			str = data
			//fmt.Printf("map[%s]为string类型，value = %v\n", key, value)
			fmt.Printf("map[%s]为string类型，value = %v\n", key, str)
		case bool:
			fmt.Printf("map[%s]为bool类型，value = %v\n", key, value)
		case float64:
			fmt.Printf("map[%s]为float64类型，value = %v\n", key, value)
		case []string:
			fmt.Printf("map[%s]为[]string类型，value = %v\n", key, data)
		case []interface{}://"subjects"对应的数组同理会被转为[]interface{}类型，应该判断是否是[]interface{}
			fmt.Printf("map[%s]为[]interface{}类型，value = %v\n", key, data)
		}
	}
	fmt.Println("str =", str)
}
