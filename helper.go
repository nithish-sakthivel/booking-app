package main

import (
	"fmt"
	"strings"
	"time"
)

func validateUserInput(firstName string, lastName string, emailId string, noOfTickets uint, availableTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailId, "@")
	isValidTicketNumber := noOfTickets > 0 && noOfTickets <= availableTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func sendTicket(noOfTickets uint, firstName string, lastName string, emailId string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets %v %v", noOfTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending  %v to email address %v\n", ticket, emailId)
	fmt.Println("##############")
	wg.Done()
}
