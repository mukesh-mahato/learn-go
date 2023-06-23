package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; no suber or parent\

	myInfo := User{"Mukesh", "mm@gmail.com", true, 19}
	fmt.Println(myInfo)
	fmt.Printf("%+v\n", myInfo)
	fmt.Printf("Name: %v, email: %v\n", myInfo.Name, myInfo.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
