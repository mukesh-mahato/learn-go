package main

import "fmt"

func main() {
	fmt.Println("Arrays in go lang") //Arrays in go lang

	var fruitList [5]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "pineapple"

	fmt.Println("fruit list: ", fruitList)      //fruit list:  [apple mango  pineapple ]
	fmt.Println("fruit list: ", len(fruitList)) //fruit list:  5

	var vegList = [4]string{"potato", "tomato", "beans"}

	fmt.Println("veg list: ", vegList) //veg list:  [potato tomato beans ]

}
