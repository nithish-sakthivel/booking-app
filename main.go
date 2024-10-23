package main

import (
	"fmt"
	"sync"
)

var conferenceName = "Go Conference"

const totalTickets = 50

var availableTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, emailId, noOfTickets := gettingUserInput()

	//Making sure the user is giving the valid data
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailId, noOfTickets, availableTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(noOfTickets, firstName, lastName)

		wg.Add(1)
		go sendTicket(noOfTickets, firstName, lastName, emailId)

		//Call the func to print the first names
		firstNames := getFirstName()
		fmt.Printf("The first names of the bookings are %v\n", firstNames)

		if availableTickets == 0 {
			fmt.Println("Sorry, all the tickets are sold out. Come back to the next one.")

		}
	} else {
		if !isValidName {
			fmt.Println("The First name or Last name you entered is too short: minimum length=2")
		}
		if !isValidEmail {
			fmt.Println("Check that you entered the correct email address")
		}
		if !isValidTicketNumber {
			fmt.Println("Please enter the correct ticket number you wanted!!")
		}
	}
	wg.Wait()
}

//Function to greet users

func greetUsers() {

	fmt.Printf("Hello everyone, welcome to %v\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still available grab your tickets soon\n", totalTickets, availableTickets)
}

// Func to print the frst name
func getFirstName() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func gettingUserInput() (string, string, string, uint) {
	// ASKING THE USER DETAILS
	var firstName string
	var lastName string
	var emailId string

	fmt.Println("Enter your First Name")
	fmt.Scanf("%s", &firstName)
	fmt.Println("Enter your Last Name")
	fmt.Scanf("%s", &lastName)
	fmt.Println("Enter your EmailID")
	fmt.Scanf("%s", &emailId)

	//Asking user for no of tickets wanted
	var noOfTickets uint

	fmt.Printf("Enter the no of tickets you wanted\n")
	fmt.Scan(&noOfTickets)
	return firstName, lastName, emailId, noOfTickets

}

func bookTicket(noOfTickets uint, firstName string, lastName string) {
	//Updating the available tickets after user booked
	availableTickets = availableTickets - noOfTickets

	//Creating map for the user data
	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		noOfTickets: noOfTickets,
	}

	//Adding the booked users to the slices to see who are all booked and keep track of it
	bookings = append(bookings, userData)
	//Display to the user on booking
	fmt.Printf("Hey %v %v thanks for choosing %v, Your %v tickets were successfully booked, Ticket was shared to your email.\n", firstName, lastName, conferenceName, noOfTickets)
	fmt.Printf("There are still %v tickets available for %v\n", availableTickets, conferenceName)

}
