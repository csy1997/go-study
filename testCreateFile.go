package main

import (
	"fmt"
	"os"
)

func main() {

	url := "day8/temp/1.txt"
	os.MkdirAll("day8/temp", os.ModePerm)
	file, err := os.Create(url)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("file =", file.Name())
}
