package main

import "fmt"

//function that switches temps
func switchTemp(fTemp float64) float64 {
	return ((fTemp - 32) * 5 / 9)
}

func main() {
	var fTemp float64 = 100

	fmt.Println(switchTemp(fTemp), "Celsius")
}
