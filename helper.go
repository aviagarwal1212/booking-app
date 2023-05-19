package main

import (
	"strings"
)

func validateUserInput(userInput userData) (bool, bool, bool) {
	isValidName := len(userInput.firstName) >= 2 && len(userInput.firstName) <= 10 && len(userInput.lastName) <= 15
	isValidEmail := strings.Contains(userInput.email, "@")
	isValidTicketNumber := userInput.userTickets > 0 && userInput.userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
