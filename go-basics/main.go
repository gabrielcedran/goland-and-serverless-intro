package main

import (
	"fmt"
	"goBasics/tickets"
)

func main() {
	fmt.Println("hello world!")

	newTicket := tickets.Ticket{
		ID:    123,
		Event: "Musical",
	}
	newTicket.PrintEvent()

}
