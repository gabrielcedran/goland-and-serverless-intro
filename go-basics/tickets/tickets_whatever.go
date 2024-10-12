package tickets

import "fmt"

type Ticket struct {
	ID    int
	Event string
}

func (t Ticket) PrintEvent() {
	fmt.Println(t.Event)
}
