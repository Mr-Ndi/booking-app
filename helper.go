package main

import "strings"

func ValidateUserInput(firstName string, lastName string, userTicket uint, email string, RemainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 4 && len(lastName) >= 2
	isValiEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= RemainingTicket

	return isValidName, isValiEmail, isValidTicketNumber
}
