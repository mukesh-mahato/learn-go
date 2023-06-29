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

	data := []int{1, 2, 3, 4, 5}
	processLargeData(&data)
	fmt.Println(data)

	var a int = 42
	var b *int = &a
	fmt.Println(a, *b) //42 42
	a = 25
	fmt.Println(a, *b) //25 25
	*b = 15
	fmt.Println(a, *b) //15 15

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Printf("%v %p %p\n", x, y, z) //[1 2 3] 0xc00001c1f8 0xc00001c200
}

/* Efficiently passing large data structures: Pointers are commonly used to avoid unnecessary copying of large data structures when passing them as function arguments. Instead of passing the entire data structure, you can pass a pointer to it, allowing functions to directly modify the original data.*/

func processLargeData(data *[]int) {
	// Modify data directly
	(*data)[0] = 10
}
