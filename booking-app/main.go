package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var conferenceName = "Go Conference" // same var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTicket uint = 50 //uint only positive value can be stored

// var bookings [50]string       //var bookings =[50]string{"ammar","tom","nana"}
// var bookings = make([]map[string]string, 0) using map

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers() // this Greet user and tell about the remaining ticket
	//fmt.Printf("conferenceName is %T, reamaingTicket is %T, conferenceTickets is %T\n", conferenceName, remainingTicket, conferenceTickets)

	for { // for remainingTicket >0 && len(bookings) <50

		firstName, lastName, email, userTickets := getUserInput()
		// ^ Take user input

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket)
		// ^ check Whether the input is valid or not

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, userTickets, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName() // this function print first name of user

			fmt.Printf("The first name of user is %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("Our conference end here")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("No of ticket you entered is in valid")
			}

		}
	}
	wg.Wait()

}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v tickets and %v are reamaining.\n", conferenceTickets, remainingTicket)

}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName) // booking["first name"] using element from map
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter numbers of ticket")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	remainingTicket = remainingTicket - uint(userTickets)

	// create a map
	/*	var userData = make(map[string]string) using map
		userData["first Name"] = firstName
		userData["last Name"] = lastName
		userData["email"] = email
		userData["Number of Tickets"] = strconv.FormatUint(uint64(userTickets), 10) */

	// using struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//bookings[0] = firstName + " " + lastName  OR
	bookings = append(bookings, userData)
	fmt.Printf("List of booking %v\n", bookings)
	/* fmt.Printf("the whole array %v\n", bookings)
	fmt.Printf("The first value %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array length %v\n", len(bookings)) */

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v ticket reamaing for %v\n", remainingTicket, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()

}
