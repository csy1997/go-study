package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "go", 3: "yoyo"}
	fmt.Println(m)

	delete(m, 1)
	fmt.Println(m)
}
