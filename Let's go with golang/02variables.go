package main

import "fmt"

const LoginToken string = "fdaggbsdbafmdb" //Public

func main() {
	var username string = "knoxartiste"
	fmt.Println(username)                              //knoxartiste
	fmt.Printf("variable is of type: %T \n", username) //variable is of type: string

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)                              //true
	fmt.Printf("variable is of type: %T \n", isLoggedIn) //variable is of type: bool

	var smallVal uint8 = 255
	fmt.Println(smallVal)                              //225
	fmt.Printf("variable is of type: %T \n", smallVal) //variable is of type: uint8

	var smallFloat float32 = 255.4545674746
	fmt.Println(smallFloat)                              //255.45457
	fmt.Printf("variable is of type: %T \n", smallFloat) //variable is of type: float32

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)                              //0
	fmt.Printf("Variable is of type: %T \n", anotherVariable) //Variable is of type: int

	// implicit type
	var website = "go.dev"
	fmt.Println(website) //go.dev

	// no var style
	numberOfUser := 3000
	fmt.Println(numberOfUser) //3000

	fmt.Println(LoginToken)                              //fdaggbsdbafmdb
	fmt.Printf("Variable is of type: %T \n", LoginToken) //Variable is of type: string

}
