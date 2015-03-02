package main

import "fmt"

func minAndMax(a []float64) {
	min := a[0]
	max := a[0]

	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
		if min > a[i] {
			min = a[i]
		}
	}

	fmt.Println("Max:", max, "Min:", min)
}
func main() {
	mySlice := []float64{10, 40, 30, 20}
	minAndMax(mySlice)
}
