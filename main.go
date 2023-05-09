package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	remainingTickets := conferenceTickets

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have %v of %v tickets remaining. Get your tickets here to attend.\n\n", remainingTickets, conferenceTickets)

	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want to buy: ")
		fmt.Scan(&userTickets)

		fmt.Println()

		isValidName := len(firstName) >= 2 && len(firstName) <= 10 && len(lastName) <= 15
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation email at %v.\nThere are %v tickets remaining.\n", firstName, lastName, userTickets, email, remainingTickets)

			firstNames := []string{}
			for _, fullName := range bookings {
				var names = strings.Fields(fullName)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all the first names of the bookings: %v\n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name entered is invalid.")
			}
			if !isValidEmail {
				fmt.Println("Email entered is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid.")
			}
		}

	}
}
