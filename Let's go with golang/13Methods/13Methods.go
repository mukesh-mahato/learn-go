package main

import "fmt"

func main() {
	fmt.Println("Methods in golang") //Methods in golang

	// no inheritance in golang; no suber or parent\

	myInfo := User{"Mukesh", "mm@gmail.com", true, 19}
	fmt.Println(myInfo)                                            ///{Mukesh mm@gmail.com true 19}
	fmt.Printf("%+v\n", myInfo)                                    ////{Name:Mukesh Email:mm@gmail.com Status:true Age:19}
	fmt.Printf("Name: %v, email: %v\n", myInfo.Name, myInfo.Email) //Name: Mukesh, email: mm@gmail.com
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
	fmt.Println("Is user active: ", u.Status) //Is user active:  true
}

func (u User) NewMail() {
	u.Email = "mukesh@gmail.com"
	fmt.Println("New email of this user is: ", u.Email) //New email of this user is:  mukesh@gmail.com
}
