package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // same var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTicket uint = 50 //uint only positive value can be stored
	// var bookings [50]string       //var bookings =[50]string{"ammar","tom","nana"}
	var bookings []string

	fmt.Printf("conferenceName is %T, reamaingTicket is %T, conferenceTickets is %T\n", conferenceName, remainingTicket, conferenceTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are reamaining.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	for { // for remainingTicket >0 && len(bookings) <50
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("Enter numbers of ticket")
		fmt.Scan(&userTickets)

		if userTickets > int(remainingTicket) {
			fmt.Printf("we only have %v ticket reaming, so you can't book %v tickets\n", remainingTicket, userTickets)
			continue
		}

		remainingTicket = remainingTicket - uint(userTickets)
		//bookings[0] = firstName + " " + lastName  OR
		bookings = append(bookings, firstName+" "+lastName)
		/* fmt.Printf("the whole array %v\n", bookings)
		fmt.Printf("The first value %v\n", bookings[0])
		fmt.Printf("Array type %T\n", bookings)
		fmt.Printf("Array length %v\n", len(bookings)) */

		fmt.Printf("Thank you %v %v for booking %v tickets.You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v ticket reamaing for %v", remainingTicket, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		noTicketRemaining := remainingTicket == 0 // boolean type variable
		if noTicketRemaining {
			fmt.Println("Our conference end here")
			break
		}
	}
}
