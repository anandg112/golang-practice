package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{firstName: "Alex",
		lastName: "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 10021,
		},
	}
	alex.updateName("alexie")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
