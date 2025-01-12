package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"
var conferenceTicket = 50
var RemainingTicket uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	email        string
	ticketNumber uint
}

var wtg = sync.WaitGroup{}

func main() {

	greetUsers()
	//This is for printing the datatypes
	// fmt.Printf("conferenceName is \t:%T, RemainingTicket is\t: %T ,and conferenceTicket is \t:%T\n", conferenceName, RemainingTicket, conferenceTicket)
	// var bookings [50]string

	firstName, lastName, email, userTicket := getUserInput()

	//calling validating function
	isValidName, isValiEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, userTicket, email, RemainingTicket)

	if isValidName && isValiEmail && isValidTicketNumber {

		wtg.Add(1)
		//calling a ticket booking parameter
		bookTicket(userTicket, firstName, lastName, email)
		go sendTicket(userTicket, firstName, lastName, email)

		// calling firstname printing function
		firstNames := GetFirstNames()
		fmt.Printf("\nThe names of those who booked are :%v", firstNames)
		fmt.Printf("\n--------------------------------------------------------------\n")
		noTicketremaining := RemainingTicket == 0

		if noTicketremaining {
			fmt.Print("The whole conference is booked out. Come back next year")
			// break
		}
	} else {
		fmt.Printf("\n--------------------------------------------------------------\n")
		// fmt.Printf("Sorry we only have %v tickets, so u can't get %v tickets, So let's start over", RemainingTicket, userTicket)
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
	wtg.Wait()

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
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTicket, RemainingTicket)
	fmt.Printf("Get your ticket here and attend  :)\n")

	fmt.Printf("\n--------------------------------------------------------------\n")

}

func GetFirstNames() []string {
	firstNames := []string{}

	// for index, booking := range bookings {
	for _, booking := range bookings { // blank identifier
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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
	fmt.Printf("%v %v Thank you for booking %v tickets.\nSoon you will receive a confirmation email at %v\n", lastName, firstName, userTicket, email)
	RemainingTicket = RemainingTicket - userTicket
	// bookings[0] = firstName + " " + lastName
	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		ticketNumber: userTicket,
	}
	bookings = append(bookings, userData)
	fmt.Printf("\nThe remaining tickets %v\n", RemainingTicket)
	fmt.Printf("\n--------------------------------------------------------------\n")

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	fmt.Printf("\n**************************************************************\n")
	fmt.Println("Wait we are sending ticket ...")
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v,", userTicket, firstName, lastName)
	fmt.Printf("\n**************************************************************\n")
	fmt.Printf("\n %v\n To email address %v\n", ticket, email)
	fmt.Printf("\n**************************************************************\n")
	wtg.Done()
}
