package main

import "fmt"

//makes a grouping of similar datatypes
type car struct {
	color string
	cost  int
}

//const cannot be changed and can be used everywhere
const (
	lots int = 5
)

//passes in an array and uses and lamda function to print length
func viewCar(a [10]car, printLots func(a [10]car)) {
	for i := 0; i < lots; i++ {
		fmt.Println(a[i].color, a[i].cost)
	}
	printLots(a)
}

//prints a length of the array to modulize
//length will get length of string or array
//length does not take value of an integer
func viewLots(a [10]car) {
	fmt.Println(len(a))
}

//This should return a sorted value without changing values in main
//returns an integer from const as well
func sortCarValue(a [10]car) ([10]car, int) {
	var temp car
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if a[i].cost > a[j].cost {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a, lots
}

//This function changes the value in main as well
//returns an integer from const as well
func sortCarReference(a *[10]car) (*[10]car, int) {
	var temp car
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if a[i].cost > a[j].cost {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a, lots
}

func main() {
	//creates an array that can hold 10 cars
	var myNum = 1
	var a [10]car
	var a1 = 10
	//infers assignment of the memory address
	var c = &a1
	//set arrays with cars and set their attributes as well
	a[0] = car{"Red", 1500}
	a[1] = car{"Blue", 2500}
	a[2] = car{"Green", 750}
	a[3] = car{"Black", 1250}
	a[4] = car{"White", 4500}
	a[5] = car{"Pink", 650}
	//print out memory adress as well as cars attributes
	fmt.Println("Hi there I am the memory value!:", c)
	fmt.Println("And I am the value!:", *c)
	//loop print on ordered cars
	fmt.Println("\nOriginal lot of Cars:")
	viewCar(a, viewLots)
	//shows the length of the array depending on switch
	if myNum == 0 {
		fmt.Println(len(a))
	} else {
		myNum = 0
	}
	switch myNum {
	case 0:
		fmt.Println("Case Zero")
	case 1:
		fmt.Println("Case One")
	default:
		fmt.Println("No Way")
	}

	//sort car by least to most by value, returns sorted value
	fmt.Println("\nCars sorted by value:")
	var b, _ = sortCarValue(a)
	viewCar(b, viewLots)
	//sort car by reference
	fmt.Println("\nCars sorted by Reference:")
	sortCarReference(&a)
	viewCar(a, viewLots)
}
