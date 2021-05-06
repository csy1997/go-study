package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (res int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		res = a / b
	}
	return
}

func main() {
	res, err := MyDiv(10, 0)
	fmt.Printf("res = %d, err = %v\n", res, err)

	res, err = MyDiv(10, 1)
	fmt.Printf("res = %d, err = %v\n", res, err)
}
