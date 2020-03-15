package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@mail.com",
			zipCode: 160036,
		},
	}

	jim.updateFirstName("Jimbo")
	jim.print()

}

func (p *person) updateFirstName(f string) {
	p.firstName = f
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
