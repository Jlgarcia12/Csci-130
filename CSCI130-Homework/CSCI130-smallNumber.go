package main

import "fmt"

func findSmall(a []int) {
	sNum := a[0]
	for i := 0; i < len(a); i++ {
		if sNum > a[i] {
			sNum = a[i]
		}
		fmt.Println(sNum, "Steps")
	}
	fmt.Println(sNum, "Final Number")
}

func main() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(x)
	findSmall(x)

}
