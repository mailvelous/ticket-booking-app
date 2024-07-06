package main

import (
	"fmt"

)

func main() {
	var concertName = "L'Arc en Ciel - The 2nd Edition"

	const numberOfTickets = 50
	var remainingTickets = numberOfTickets
	var bookings = []string{}


	fmt.Println("Welcome to", concertName, "Ticket Booking App")
	fmt.Println("Total Tickets :", numberOfTickets, "Remaining Tickets :", remainingTickets)
	fmt.Println("--- Get your Ticket NOW!!! ---")

	var userName string
	var userTickets int
	var userEmail string


	fmt.Println("Enter your name :")
	fmt.Scan(&userName)

	fmt.Println("Enter your email address :")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets you want?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName)

	fmt.Println("Thank you ", userName, " for booking", userTickets, "tickets. You will be receiving a confirmation email at", userEmail )
	fmt.Println("Remaining Tickets :", remainingTickets)

	fmt.Println("These are all our bookings: ", bookings, "has booked", userTickets, "tickets")
}