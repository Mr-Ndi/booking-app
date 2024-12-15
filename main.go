package main

import "fmt"

func main() {
	var ConferenceName = "Go conference"
	const conferenceTicket = 50
	var remainingTicket = 50

	fmt.Println("Wellcome to", ConferenceName, " booking application !")
	fmt.Println("We have a total of", conferenceTicket, "and", remainingTicket, "are still available")
	fmt.Println("Get your ticket here and attend  :)")

}
