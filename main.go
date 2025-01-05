package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go conference"
var conferenceTicket = 50
var remainingTicket uint = 50
var bookings []string

func main() {

	greetUsers()
	//This is for printing the datatypes
	// fmt.Printf("conferenceName is \t:%T, remainingTicket is\t: %T ,and conferenceTicket is \t:%T\n", conferenceName, remainingTicket, conferenceTicket)
	// var bookings [50]string

	for {
		firstName, lastName, email, userTicket := getUserInput()
		//calling validating function
		isValidName, isValiEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, userTicket, email)

		if isValidName && isValiEmail && isValidTicketNumber {
			//calling a ticket booking parameter
			bookTicket(userTicket, firstName, lastName, email)
			// calling firstname printing function
			firstNames := GetFirstNames()
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

func greetUsers() {

	fmt.Printf("\n--------------------------------------------------------------\n")
	fmt.Printf("Wellcome to %v booking application !\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here and attend  :)\n")

	fmt.Printf("\n--------------------------------------------------------------\n")

}

func GetFirstNames() []string {
	firstNames := []string{}

	// for index, booking := range bookings {
	for _, booking := range bookings { // blank identifier
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	// fmt.Printf("\nThe names of those who booked are :%v", firstNames)
	return firstNames
}

func ValidateUserInput(firstName string, lastName string, userTicket uint, email string) (bool, bool, bool) {
	isValidName := len(firstName) >= 4 && len(lastName) >= 2
	isValiEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	return isValidName, isValiEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	fmt.Printf("\n--------------------------------------------------------------\n")
	fmt.Printf("%v %v Thank you for booking %v tickets. You will receive a confirmation email at %v\n", lastName, firstName, userTicket, email)
	remainingTicket = remainingTicket - userTicket
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("\nThe remaining tickets %v\n", remainingTicket)
	fmt.Printf("\n--------------------------------------------------------------\n")

}
