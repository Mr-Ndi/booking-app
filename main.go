package main

import (
	"fmt"
	"strings"
)

func main() {
	ConferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50

	//This is for printing the datatypes
	// fmt.Printf("ConferenceName is \t:%T, remainingTicket is\t: %T ,and conferenceTicket is \t:%T\n", ConferenceName, remainingTicket, conferenceTicket)

	fmt.Printf("\n--------------------------------------------------------------\n")
	fmt.Printf("Wellcome to %v booking application !\n", ConferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here and attend  :)\n")

	fmt.Printf("\n--------------------------------------------------------------\n")

	// var bookings [50]string
	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("\nHey what's your first name pls :)\t:")
		fmt.Scan(&firstName)

		fmt.Println("Then what's your second name pls :)\t:")
		fmt.Scan(&lastName)

		fmt.Printf("Lastly %v can you tell me how many ticket would you like to book\t:", firstName)
		fmt.Scan(&userTicket)

		fmt.Println("Then what's your email address :)\t:")
		fmt.Scan(&email)

		fmt.Printf("\n--------------------------------------------------------------\n")
		fmt.Printf("%v %v Thank you for booking %v tickets. You will receive a confirmation email at %v\n", lastName, firstName, userTicket, email)
		remainingTicket = remainingTicket - userTicket
		fmt.Printf("\nThe remaining tickets %v\n", remainingTicket)
		fmt.Printf("\n--------------------------------------------------------------\n")

		bookings = append(bookings, firstName+" "+lastName)
		firstNames := []string{}
		// for index, booking := range bookings {
		for _, booking := range bookings { // blank identifier
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("\nThe names of those who booked are :%v", firstNames)
		// bookings[0] = firstName + " " + lastName

	}

	// fmt.Printf("The whole array %v\n", bookings)
	// fmt.Printf("The first Element %v\n", bookings[0])
	// fmt.Printf("The array length %v\n", len(bookings))
	// fmt.Printf("The array Type %T", bookings)
}
