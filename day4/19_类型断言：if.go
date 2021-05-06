package main

import "fmt"

type Student15 struct {
	name string
	id int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "abc"
	i[2] = Student15{"mike", 666}

	//类型查询（断言）
	for index, data := range i {
		//第一个返回的是值，第二个返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("s[%d]类型为int，内容为%v\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("s[%d]类型为string，内容为%v\n", index, value)
		} else if value, ok := data.(Student15); ok == true {
			fmt.Printf("s[%d]类型为Student15，内容为%v\n", index, value)
		}
	}
}
