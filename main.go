package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets = conferenceTickets
var bookings = []userData{}

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		userInput := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userInput)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = bookTicket(userInput)
			printNames()

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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v of %v tickets remaining. Get your tickets here to attend.\n\n", remainingTickets, conferenceTickets)
}

func printNames() {
	fmt.Printf("These are all the first names of the bookings: %v\n\n", bookings)
}

func getUserInput() userData {
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

	userInput := userData{firstName, lastName, email, userTickets}

	return userInput
}

func bookTicket(userInput userData) uint {
	remainingTickets := remainingTickets - userInput.userTickets
	bookings = append(bookings, userInput)
	fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation email at %v.\nThere are %v tickets remaining.\n\n", userInput.firstName, userInput.lastName, userInput.userTickets, userInput.email, remainingTickets)
	wg.Add(1)
	go sendTickets(userInput)
	return remainingTickets
}

func sendTickets(userInput userData) {
	time.Sleep(10 * time.Second)
	fmt.Println()
	fmt.Printf("EMAILED: %v tickets to %v %v at %v\n", userInput.userTickets, userInput.firstName, userInput.lastName, userInput.email)
	fmt.Println()
	wg.Done()
}
