package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTicket int = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidatUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstnames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end programme

				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered is invalid, does not contain @ sign ")

			}
			if !isValidTicketNumber {
				fmt.Println(" Number of ticket you entered is invalid")
			}
		}
	}
}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application: \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstnames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email name: ")
	fmt.Scan(&email)
	fmt.Printf("Enter your tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}