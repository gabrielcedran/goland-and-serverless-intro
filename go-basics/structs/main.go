package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{
		Name: name,
		Age:  age,
	}
	fmt.Printf("Memory address inside func %p\n", &p.Name)
	return &p
}

func changeAgeCopy(person Person, age int) {
	fmt.Println("Allocated memory of age in copy", &person.Age)
	person.Age = age
}

func (p Person) changeAgeReceiver(age int) {
	p.Age = age
}

func (p *Person) changeAgeReceiverReference(age int) {
	p.Age = age
	fmt.Println("Allocated memory of age in ref", &p.Age)
}

func main() {

	myPerson := NewPerson("Don Bob", 9)
	fmt.Printf("Memory address inside func %p\n", &myPerson.Name)
	fmt.Println("Allocated memory of age in main", &myPerson.Age)

	changeAgeCopy(*myPerson, 10)
	fmt.Printf("This is my person %+v\n", myPerson)

	myPerson.changeAgeReceiver(11)
	fmt.Printf("This is my person %+v\n", myPerson)

	myPerson.changeAgeReceiverReference(12)
	fmt.Printf("This is my person %+v\n", myPerson)

}
