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

    	validName := len(userFirstName) >= 2 && len(userLastName) >= 2
    	validEmail := strings.Contains(userEmail, "@")
    	validTickets := userTickets > 0 && userTickets <= remainingTickets

		if validEmail && validTickets && validName{	
			remainingTickets -= userTickets
			bookings = append(bookings, userFirstName+" "+userLastName)

			if userTickets == 1 {
				fmt.Printf("Thanks %v %v  for order %v ticket, we will contact from this email %v\n", userFirstName, userLastName, userTickets, userEmail)
			} else {
				fmt.Printf("Thanks %v %v for order %v tickets, we will contact from this email %v\n", userFirstName, userLastName, userTickets, userEmail)
			}
			fmt.Printf("We have %v tickets remaing\n", remainingTickets) 
		fmt.Printf("You are add to the list of bookings, this is the total members %v\n", len(bookings))
		firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
				fmt.Printf("The list of all the members fist names is %v\n", firstNames)
			}
		}else{
			if !validEmail{
				fmt.Println("Sorry but yout email must contain @ carachter")
				fmt.Println("What is your email?")
				fmt.Scan(&userEmail)
			}else if !validTickets{
				if remainingTickets == 1{
					fmt.Printf("Sorry we only have %v ticket", remainingTickets)
					fmt.Println("How many tickets you want?")
					fmt.Scan(&userTickets)
				}else{
					fmt.Printf("Sorry we only have %v tickets", remainingTickets)
					fmt.Println("How many tickets you want?")
					fmt.Scan(&userTickets)
				}
			}else if !validName{
				fmt.Println("Sorry but your name need to have a minimum of 2 characters")

				fmt.Println("What is your first name?")
				fmt.Scan(&userFirstName)
		
				fmt.Println("What is your last name?")
				fmt.Scan(&userLastName)
			}
		}
	}
}
