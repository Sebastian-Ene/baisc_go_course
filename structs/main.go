package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email       string
	phoneNumber int
}

func main() {
	pers1 := person{
		firstName: "Sebi",
		lastName:  "Yeah",
		contact: contact{
			email:       "something@asd.com",
			phoneNumber: 1235,
		},
	}
	pers1.updateName("Johnie")
	pers1.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
	p.print()
}
