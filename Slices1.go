package main

import "fmt"

//makes a type of string
type stringMethod string

//method that has an input and function
func (s stringMethod) exclaim() {
	fmt.Println(s + "!")
}

//given any length of array will print to console
func printName(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func main() {
	var name = "Jose"
	//needs to be defined as a type in order to use methods
	var name1 stringMethod = "Epi"
	mySlice := []string{"Jose", "Brandon", "John"}
	//prints what is given
	printName(mySlice[1:2])
	//nets you the first and second element in slice
	fmt.Println(mySlice[:1])
	//gets the second char in the string
	fmt.Println(name[1:2])
	//appends adds onto the slice and set it to mySlice
	mySlice = append(mySlice, mySlice[1:]...)
	//prints out mySlice
	fmt.Println(mySlice)
	//appends slice and var into it
	mySlice = append(mySlice[2:3], name)
	//prints changed slice
	fmt.Println(mySlice)
	//uses type and calls function in order to rund the func
	name1.exclaim()
}
