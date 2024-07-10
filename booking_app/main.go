package main

import (
	"fmt"
	"strings"
)

func main(){
	eventname := "digital gandhi"
	const conferenceTickets=350
	var remainingtickets uint=350

	fmt.Printf("welcome to %v powered by \n  ", eventname)
	fmt.Println("we have a total of", conferenceTickets, "tickets and", remainingtickets,"are still available")
	fmt.Println("get your tickets right now ")

	var bookings =[]string{"nicole", "peter"}
	
	var firstname string 
	var lastname string
	var email string 
	var userticket int 


	for {
	fmt.Println(("enter your first name"))
	fmt.Scan(&firstname)
	fmt.Println(&firstname)
	fmt.Println(("enter your last name"))
	fmt.Scan(&lastname)
	fmt.Println(("enter your email "))
	fmt.Scan(&email)
	fmt.Println(("enter number of tickets "))
	fmt.Scan(&userticket)

	remainingtickets= remainingtickets - uint(userticket)
	bookings=append(bookings, firstname +" "+lastname)
	println("array length", len(bookings))
	fmt.Printf("thank you %v %v for booking %v tickets for %v event, you will receive a confirmation email at %v ",firstname, lastname, userticket, eventname, email)	
	fmt.Printf("remaining tickets %v", remainingtickets)

	firstnames := []string{}
	for _, val := range bookings{
		var fname=strings.Fields(val)
		
		firstnames =append(firstnames, fname[0])

	} //index and value in bookings in range
	fmt.Printf("the first names of bookings are %v \n",firstnames)
	if remainingtickets==0 {
		fmt.Println("event is booked out")
		break
	}
	}
	
}


