package main

import (
	"fmt"
	"goBasics/circle"
	"goBasics/tickets"
)

func main() {
	fmt.Println("hello world!")

	newTicket := tickets.Ticket{
		ID:    123,
		Event: "Musical",
	}
	newTicket.PrintEvent()

	myCircle := circle.Circle{Radius: 12.0}
	myCircle.CalculcateCircumference()
	fmt.Println(myCircle.Circumference)
	myCircle.Area()
}
