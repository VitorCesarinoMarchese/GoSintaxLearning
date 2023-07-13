package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Porra"
	const tickets uint8 = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// fmt.Printf("Tickets is %T, Name is %T /n", tickets, name)
	fmt.Printf("Welcome to my %v\n", name)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", tickets, remainingTickets)
	fmt.Println("Buy your tickets here TA!")

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

		if int(remainingTickets)-int(userTickets) < 0 {
			for int(remainingTickets)-int(userTickets) < 0 {
				if remainingTickets == 1 {
					fmt.Printf("Sorry but we only have %v ticket\n", remainingTickets)
				} else {
					fmt.Printf("Sorry but we only have %v tickets\n", remainingTickets)
				}
				fmt.Println("How many tickets you want?")
				fmt.Scan(&userTickets)
			}
		}

		remainingTickets -= userTickets
		bookings = append(bookings, userFirstName+" "+userLastName)

		if userTickets == 1 {
			fmt.Printf("Thanks %v %v  for order %v ticket, we will contact from this email %v\n", userFirstName, userLastName, userTickets, userEmail)
		} else {
			fmt.Printf("Thanks %v %v for order %v tickets, we will contact from this email %v\n", userFirstName, userLastName, userTickets, userEmail)
		}
		fmt.Printf("We have %v tickets remaing\n", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
			fmt.Printf("The list of all the members fist names is %v\n", firstNames)
		}
		fmt.Printf("You are add to the list of bookings, this is the total members %v\n", len(bookings))
	}
}
