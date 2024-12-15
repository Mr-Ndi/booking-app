package main

import "fmt"

func main() {
	var ConferenceName = "Go conference"
	const conferenceTicket = 50
	var remainingTicket = 50

	fmt.Printf("Wellcome to %v booking application !\n", ConferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here and attend  :)")

}
