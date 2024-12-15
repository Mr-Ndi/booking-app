package main

import "fmt"

func main() {
	ConferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket = 50

	//This is for printing the datatypes
	// fmt.Printf("ConferenceName is \t:%T, remainingTicket is\t: %T ,and conferenceTicket is \t:%T\n", ConferenceName, remainingTicket, conferenceTicket)

	fmt.Printf("Wellcome to %v booking application !\n", ConferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here and attend  :)\n")

	fmt.Printf("\n--------------------------------------------------------------\n")

	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Printf("Then what's your first name pls :)\t:")
	fmt.Scan(&firstName)

	fmt.Printf("Then what's your first name pls :)\t:")
	fmt.Scan(&lastName)

	fmt.Printf("Then %v how many ticket you have booked\t:", firstName)
	fmt.Scan(&userTicket)

	fmt.Printf("Then what's your email address :)\t:")
	fmt.Scan(&email)
	// userName = "Serge"
	// userTicket = 2

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
}
