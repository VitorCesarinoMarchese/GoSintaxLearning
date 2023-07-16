package main

import (
	"fmt"
	"strings"
	"time"
)

var name = "Porra"

const tickets uint8 = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	tickets uint
}

func main() {

	greetUsers()

	for remainingTickets != 0 {

		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets uint

		fmt.Println("What is your first name?")
		fmt.Scan(&userFirstName)

		fmt.Println("What is your last name?")
		fmt.Scan(&userLastName)

		fmt.Println("What is your email?")
		fmt.Scan(&userEmail)

		fmt.Println("How many tickets you want?")
		fmt.Scan(&userTickets)

		if len(userFirstName) < 2 && len(userLastName) < 2 {
			for len(userFirstName) < 2 && len(userLastName) < 2 {
				fmt.Println("Sorry but your first name or last name needs to contain at least 2 characters")

				fmt.Println("What is your first name?")
				fmt.Scan(&userFirstName)

				fmt.Println("What is your last name?")
				fmt.Scan(&userLastName)
			}
		}
		if userTickets < 0 || remainingTickets < userTickets {
			for userTickets < 0 || remainingTickets < userTickets {
				fmt.Printf("Sorry but we only have %v tickets\n", remainingTickets)

				fmt.Println("How many tickets you want?")
				fmt.Scan(&userTickets)
			}
		}
		if !strings.Contains(userEmail, "@") {
			for !strings.Contains(userEmail, "@") {
				fmt.Println("Sorry but your email needs to contain a @ symbol")

				fmt.Println("What is your email?")
				fmt.Scan(&userEmail)
			}
		}

		// fmt.Printf("validName: %v, validEmail: %v, validTickets %v\n", validName, validEmail, validTickets)

		remainingTickets -= userTickets

		userData := userData {
			firstName: userFirstName,
			lastName: userLastName,
			email: userEmail,
			tickets: userTickets,
		}


		bookings = append(bookings, userData)
		fmt.Println("Loading")
		time.Sleep(5 * time.Second)
		thanks(userTickets, userFirstName, userLastName, userEmail)
		firstName()
		fmt.Printf("%v\n", bookings)
	}
}

func greetUsers() {
	fmt.Printf("Welcome to my %v\n", name)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", tickets, remainingTickets)
	fmt.Println("Buy your tickets here TA!")
}
func firstName() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
		fmt.Printf("The list of all the members fist names is %v\n", firstNames)
	}
}
func thanks(userTickets uint, userFirstName string, userLastName string, userEmail string) {
	if userTickets == 1 {
		fmt.Printf("Thanks %v %v  for order %v ticket, we will contact from this email %v\n", userFirstName, userLastName, userTickets, userEmail)
	} else {
		fmt.Printf("Thanks %v %v for order %v tickets, we will contact from this email %v\n", userFirstName, userLastName, userTickets, userEmail)
	}
	fmt.Printf("We have %v tickets remaing\n", remainingTickets)
	fmt.Printf("You are add to the list of bookings, this is the total members %v\n", len(bookings))
}
