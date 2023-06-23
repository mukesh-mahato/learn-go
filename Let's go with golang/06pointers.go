package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("The value of pointer is", ptr)

	num := 98

	var ptr = &num
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is ", num)
}
