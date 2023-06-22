package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isEmailValid, isValidTicketsNumber := validUserInput(firstName, lastName, email, userTickets)

		if isValidName && isEmailValid && isValidTicketsNumber {

			bookTicket(userTickets, bookings, firstName, lastName, email)
			// bookings[0] = firstName + "" + lastName

			// fmt.Printf("The whole slice %v\n", bookings)
			// fmt.Printf("The first value : %v\n", bookings[0])
			// fmt.Printf("Slice type : %T\n", bookings)
			// fmt.Printf("Slice length : %v\n", len(bookings))

			firstNames := getFirstNames()
			fmt.Printf("The first name of booking are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Firstname or Lastname you entered is too short")
			}
			if !isEmailValid {
				fmt.Println("Email address you entered doesnot contains 'a'")
			}
			if !isValidTicketsNumber {
				fmt.Println("Ticket number you entered is invalid")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		var firstName = name[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames

}

func validUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isEmailValid, isValidTicketsNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, bookings []string, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
