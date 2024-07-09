package main

import "fmt"

func main(){
	var eventname= "Frosh"
	const conferenceTickets=350
	var remainingtickets=350

	fmt.Printf("welcome to %v powered by frosh\n  ", eventname)
	fmt.Println("we have a total of", conferenceTickets, "tickets and", remainingtickets,"are still available")
	fmt.Println("get your tickets right now ")
}

