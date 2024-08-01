package main

import (
	"fmt"
	"time"
	"sync"
)

var concertName = "L'Arc en Ciel - The 2nd Edition"

const numberOfTickets = 50

var remainingTickets = 50
var bookings = make([]UserData, 0)
var ticketBookings = []int{}

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greeting()
	firstName, lastName, userEmail, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userTickets, remainingTickets, userEmail)

	if isValidEmail && isValidName && isValidTicketNumber {

		wg.Add(1)
		bookTickets(userTickets, firstName, lastName, userEmail)

		go sendTicket(userTickets, firstName, lastName, userEmail)

		firstNames := getFirstNames()
		userTickets := getUserTickets()

		fmt.Println("These are all our bookings: ", firstNames, "has booked", userTickets)
		fmt.Printf("These are all our first names: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All tickets have been sold out.")
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid name. First and last name must be at least 2 characters long. Try again.")
		}
		if !isValidEmail {
			fmt.Println("Invalid email address. Must be contain @gmail.com. Try again.")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid number of tickets. Must be greater than 0 and less than or equal to", remainingTickets, ". Try again")
		}
	}
	wg.Wait()

	

}

func greeting() {

	fmt.Println("Welcome to the Ticket Booking App of ", concertName)
	fmt.Println("Total Tickets :", numberOfTickets, "Remaining Tickets :", remainingTickets)
	fmt.Println("--- Get your Ticket NOW!!! ---")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserTickets() []int {
	userTickets := []int{}
	for _, booking := range bookings {
		userTickets = append(userTickets, booking.userTickets)
	}
	return userTickets

}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets int

	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name :")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address :")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets you want?")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       userEmail,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	ticketBookings = append(ticketBookings, remainingTickets)

	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Println("Thank you ", firstName, " ", lastName, "for booking", userTickets, "tickets. You will be receiving a confirmation email at", userEmail)
	fmt.Println("Remaining Tickets :", remainingTickets)
}

func sendTicket(userTickets int, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("================================================================")
	fmt.Println("Sending tickets...", ticket, "to this email :", userEmail)
	fmt.Println("================================================================")
	wg.Done()

}
