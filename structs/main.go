package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "a@bz.com",
			pinCode: 2000,
		},
	}
	//alexPointer := &alex
	alex.updateName("jim")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
