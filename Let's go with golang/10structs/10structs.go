package main

import "fmt"

func main() {
	fmt.Println("Structs in golang") //Structs in golang

	// no inheritance in golang; no suber or parent\

	myInfo := User{"Mukesh", "mm@gmail.com", true, 19}
	fmt.Println(myInfo)                                            //{Mukesh mm@gmail.com true 19}
	fmt.Printf("%+v\n", myInfo)                                    //{Name:Mukesh Email:mm@gmail.com Status:true Age:19}
	fmt.Printf("Name: %v, email: %v\n", myInfo.Name, myInfo.Email) //Name: Mukesh, email: mm@gmail.com
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
