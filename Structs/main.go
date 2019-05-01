package main

import "fmt"

// embedding structs

type contactInfo struct {
	email   string
	zipCode int
}

// define struct for person(firstName, lastName)
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // embedded struct
}

func main() {
	// assign value to struct
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("Saurabh")
	jim.print()
}

// pass by value
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// struct with reciever function
func (p person) print() {
	fmt.Printf("%+v", p)
}
