package main

import "fmt"

type Student16 struct {
	name string
	id int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "abc"
	i[2] = Student16{"mike", 666}

	//类型查询（断言）
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("s[%d]类型为int，内容为%v\n", index, value)
		case string:
			fmt.Printf("s[%d]类型为string，内容为%v\n", index, value)
		case Student16:
			fmt.Printf("s[%d]类型为Student16，内容为%v\n", index, value)
		}
	}
}
