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

		fmt.Println("Then what's your email address :)\t:")
		fmt.Scan(&email)

		fmt.Printf("Lastly %v can you tell me how many ticket would you like to book\t:", firstName)
		fmt.Scan(&userTicket)

		isValidName := len(firstName) >= 4 && len(lastName) >= 2
		isValiEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

		if isValidName && isValiEmail && isValidTicketNumber {

			fmt.Printf("\n--------------------------------------------------------------\n")
			fmt.Printf("%v %v Thank you for booking %v tickets. You will receive a confirmation email at %v\n", lastName, firstName, userTicket, email)
			remainingTicket = remainingTicket - userTicket
			fmt.Printf("\nThe remaining tickets %v\n", remainingTicket)
			fmt.Printf("\n--------------------------------------------------------------\n")

			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)
			firstNames := []string{}

			// for index, booking := range bookings {
			for _, booking := range bookings { // blank identifier
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("\nThe names of those who booked are :%v", firstNames)

			noTicketremaining := remainingTicket == 0

			if noTicketremaining {
				fmt.Print("The whole conference is booked out. Come back next year")
				break
			}
		} else {
			fmt.Printf("\n--------------------------------------------------------------\n")
			// fmt.Printf("Sorry we only have %v tickets, so u can't get %v tickets, So let's start over", remainingTicket, userTicket)
			if !isValidName {
				fmt.Print("Either First name or Second name is too short")
			}
			if !isValiEmail {
				fmt.Print("The entered email isn't correct, may be  @ sign is missing")
			}
			if !isValidTicketNumber {
				fmt.Print("Hey bro checkout the ticket bumber you have entered")
			}
			fmt.Printf("\n--------------------------------------------------------------\n")
			// continue
		}

	}
	// city := "Kigali"
	// switch city {
	// 	case "Bujumbura":
	// 		//something expressable
	// 	case "Maputo", "cape-town":
	// 		//something expressible
	// 	default:
	// 		//something for undefined inputs
	// }

	// handling arrays, i mean before introduction of slices
	// fmt.Printf("The whole array %v\n", bookings)
	// fmt.Printf("The first Element %v\n", bookings[0])
	// fmt.Printf("The array length %v\n", len(bookings))
	// fmt.Printf("The array Type %T", bookings)
}
