package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in go lang")

	var fruitList = []string{"apple", "mango", "pineapple"}
	fmt.Printf("type of fruit list: %T\n", fruitList)

	fruitList = append(fruitList, "banana", "orange")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 126
	highScores[1] = 937
	highScores[2] = 646
	highScores[3] = 764
	// highScores[4] = 565

	highScores = append(highScores, 346, 243, 478)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices based on index

	var cources = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(cources)
	var index int = 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
