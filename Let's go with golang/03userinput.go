package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome) //Welcome to user input

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Pizza: ") //Enter the rating for Pizza:

	// comma ok || err ok

	input, _ := reader.ReadString(('\n'))
	fmt.Println("Thanks for rating", input)        //Thanks for rating 5
	fmt.Printf("type of this rating is %T", input) //type of this rating is string
}
