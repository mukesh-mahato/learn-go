package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; no suber or parent\

	myInfo := User{"Mukesh", "mm@gmail.com", true, 19}
	fmt.Println(myInfo)
	fmt.Printf("%+v\n", myInfo)
	fmt.Printf("Name: %v, email: %v\n", myInfo.Name, myInfo.Email)
	myInfo.GetStatus()
	myInfo.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "mukesh@gmail.com"
	fmt.Println("New email of this user is: ", u.Email)
}
