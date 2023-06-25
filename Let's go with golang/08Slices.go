package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in go lang") //veg list:  [potato tomato beans ]

	var fruitList = []string{"apple", "mango", "pineapple"}
	fmt.Printf("type of fruit list: %T\n", fruitList) //type of fruit list: []string

	fruitList = append(fruitList, "banana", "orange")
	fmt.Println(fruitList) //[apple mango pineapple banana orange]

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) //[mango pineapple banana orange]

	highScores := make([]int, 4)

	highScores[0] = 126
	highScores[1] = 937
	highScores[2] = 646
	highScores[3] = 764
	// highScores[4] = 565

	highScores = append(highScores, 346, 243, 478)

	fmt.Println(highScores) //[126 937 646 764 346 243 478]

	sort.Ints(highScores)
	fmt.Println(highScores) //[126 243 346 478 646 764 937]

	// how to remove a value from slices based on index

	var cources = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(cources) //[reactjs javascript swift python ruby]
	var index int = 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources) //[reactjs javascript python ruby]

}
