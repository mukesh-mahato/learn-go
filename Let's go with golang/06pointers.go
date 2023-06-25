package main

import "fmt"

func main() {
	fmt.Println("Pointers") //Pointers

	// var ptr *int
	// fmt.Println("The value of pointer is", ptr)

	num := 98

	var ptr = &num
	fmt.Println("Value of actual pointer is ", ptr)  //Value of actual pointer is  0xc0000180b0
	fmt.Println("Value of actual pointer is ", *ptr) //Value of actual pointer is  98

	*ptr = *ptr + 2
	fmt.Println("New value is ", num) //New value is  100
}
