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
	andi := person{
		firstName: "Andi",
		lastName:  "Noya",
		contactInfo: contactInfo{
			email:   "dulu@dulu.com",
			zipCode: 110110,
		},
	}
	fmt.Println(andi)
	var alex person
	alex.firstName = "alex"
	alex.lastName = "anderson"
	alex.contactInfo.email = "alex@blah.com"
	alex.contactInfo.zipCode = 22474
	fmt.Println(alex)
	alex.print()

	// alexPointer := &alex
	// alexPointer.updateName("Sandro")
	alex.updateName("Sandro")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
