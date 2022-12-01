package main

import "fmt"

type ContactDetails struct {
	email string
	phone int
}

type Person struct {
	firstName      string
	lastName       string
	contactDetails ContactDetails
}

func (p Person) print() {
	fmt.Println(p)
}

func (x *Person) updateName(firstName string, lastName string) {
	x.firstName = firstName
	x.lastName = lastName
}

func main() {
	p := Person{firstName: "Matheas", lastName: "Charukeas", contactDetails: ContactDetails{email: "tmatheas@gmail.com", phone: 9095558811}}
	p.print()
	p.updateName("Nageswaran", "A")
	p.print()
}
