package main

import "fmt"

func main(){
	eventname := "digital gandhi"
	const conferenceTickets=350
	var remainingtickets uint=350

	fmt.Printf("welcome to %v powered by frosh\n  ", eventname)
	fmt.Println("we have a total of", conferenceTickets, "tickets and", remainingtickets,"are still available")
	fmt.Println("get your tickets right now ")
	fmt.Scan()

	var username ="tom"
	fmt.Println((username))
}

