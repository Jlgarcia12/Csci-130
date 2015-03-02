package main

import "fmt"

func maxValue(args ...int) int {
	max := 0
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(maxValue(4, 3, 6, 7, 23, 1))
}
