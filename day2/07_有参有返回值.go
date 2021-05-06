package main

import "fmt"

func main() {
	max, min := MaxAndMin(3, 4, 2, 5, 1)
	fmt.Printf("max = %d, min = %d\n", max, min)
}

func MaxAndMin(args ...int) (max, min int) {
	max = args[0]
	min = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
		if args[i] < min {
			min = args[i]
		}
	}
	return
}
