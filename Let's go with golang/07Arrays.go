package main

import "fmt"

func main() {
	fmt.Println("Arrays in go land")

	var fruitList [5]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "pineapple"

	fmt.Println("fruit list: ", fruitList)
	fmt.Println("fruit list: ", len(fruitList))

	var vegList = [4]string{"potato", "tomato", "beans"}

	fmt.Println("veg list: ", vegList)

}
