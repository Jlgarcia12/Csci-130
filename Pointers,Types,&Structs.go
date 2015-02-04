package main

import "fmt"

//makes a grouping of similar datatypes
type car struct {
	color string
	cost  int
}

func viewCar(a [10]car) {
	for i := 0; i < 5; i++ {
		fmt.Println(a[i].color, a[i].cost)
	}
}

//This should return a sorted value without changing values in main
func sortCarValue(a [10]car) [10]car {
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
	return a
}

//This function changes the value in main as well
func sortCarReference(a *[10]car) *[10]car {
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
	return a
}

func main() {
	//creates an array that can hold 10 cars
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
	fmt.Println(" ")
	fmt.Println("Original lot of Cars:")
	viewCar(a)
	//sort car by least to most by value
	fmt.Println(" ")
	fmt.Println("Cars sorted by value:")
	sortCarValue(a)
	viewCar(a)
	//sort car by reference
	fmt.Println(" ")
	fmt.Println("Cars sorted by Reference:")
	sortCarReference(&a)
	viewCar(a)
}
