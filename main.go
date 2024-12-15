package main

import "fmt"

func main() {
	ConferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket = 50

	//This is for printing the datatypes
	fmt.Printf("ConferenceName is \t:%T, remainingTicket is\t: %T ,and conferenceTicket is \t:%T\n", ConferenceName, remainingTicket, conferenceTicket)

	fmt.Printf("Wellcome to %v booking application !\n", ConferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here and attend  :)\n")

	var userName string
	var userTicket int

	// then lets assumes that the user will enter the name and nbr of tickets
	userName = "Serge"
	userTicket = 2

	fmt.Printf("%v Has booked %v tickets\n", userName, userTicket)
}
