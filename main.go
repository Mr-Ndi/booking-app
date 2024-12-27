package main

import "fmt"

func main() {
	ConferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50

	//This is for printing the datatypes
	// fmt.Printf("ConferenceName is \t:%T, remainingTicket is\t: %T ,and conferenceTicket is \t:%T\n", ConferenceName, remainingTicket, conferenceTicket)

	fmt.Printf("Wellcome to %v booking application !\n", ConferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here and attend  :)\n")

	fmt.Printf("\n--------------------------------------------------------------\n")

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Hey what's your first name pls :)\t:")
	fmt.Scan(&firstName)

	fmt.Println("Then what's your second name pls :)\t:")
	fmt.Scan(&lastName)

	fmt.Printf("Lastly %v can you tell me how many ticket would you like to book\t:", firstName)
	fmt.Scan(&userTicket)

	remainingTicket = remainingTicket - userTicket

	fmt.Println("Then what's your email address :)\t:")
	fmt.Scan(&email)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("Thank you for booking your titckets Then the reamaining tickets %v\n", remainingTicket)
}
