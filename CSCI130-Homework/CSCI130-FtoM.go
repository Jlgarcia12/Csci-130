package main

import "fmt"

//function that switches temps
func switchDistance(fDis float64) float64 {
	return (fDis * .3048)
}

func main() {
	var fDis float64 = 10

	fmt.Println(switchDistance(fDis), "meters")
}
