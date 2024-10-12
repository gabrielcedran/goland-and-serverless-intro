package circle

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	Radius        float64
	Circumference float64
}

func (c *Circle) CalculcateCircumference() {
	c.Circumference = c.Radius * PI * 2
	fmt.Println(c.Circumference)
}

func (c *Circle) Area() {
	fmt.Println(c.Radius * c.Radius * PI)
}
