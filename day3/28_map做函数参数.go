package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "go", 3: "yoyo"}
	fmt.Println(m)

	//map是引用传递
	testDelete(m)
	fmt.Println(m)
}

func testDelete(m map[int]string) {
	delete(m, 1)
}
