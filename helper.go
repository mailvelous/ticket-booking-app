package main

import (
    "strings"
)

func validateUserInput(firstName, lastName string, userTickets int, remainingTickets int, userEmail string) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@gmail.com")
	isValidTicketNumber := userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTicketNumber
}
