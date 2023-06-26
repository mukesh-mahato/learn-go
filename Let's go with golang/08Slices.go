package main

import (
	"fmt"
	"sort"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

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

	// Dynamic lists: Slices are commonly used to represent dynamic lists of elements. They allow for easy appending, removing, and modifying elements without having to manually manage array sizes.
	numbers := []int{1, 2, 3, 4, 5}

	// Append an element
	numbers = append(numbers, 6)

	// Remove an element at index 2
	numbers = append(numbers[:2], numbers[3:]...)

	fmt.Println(numbers) // Output: [1 2 4 5 6]

	// Implementing data structures: Slices are often used as the underlying data structure for implementing various data structures such as stacks, queues, and dynamic arrays. They allow for efficient insertion, deletion, and resizing operations.

	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println(stack.Pop()) // Output: 30
	fmt.Println(stack.Pop()) // Output: 20

}
