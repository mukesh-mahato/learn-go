package main

import "fmt"

func main() {
	fmt.Println("Functions in golang") //Functions in golang

	result := add(4, 5)
	fmt.Println(result) //9

	proResult := proAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(proResult) //28
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func proAdd(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}
