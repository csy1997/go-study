package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666

	//编码成json
	//res, err := json.Marshal(m)
	res, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("res =", string(res))
	}
}
