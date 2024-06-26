package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	dat := person{
		firstname: "Dat",
		lastname:  "Nguyen",
		contactInfo: contactInfo{
			zipCode: 700000,
			email:   "datntt3@fpt.com",
		},
	}
	//huy := person{"Huy", "Nguyen"}
	var mrX person

	mrX.firstname = dat.firstname
	mrX.lastname = dat.lastname
	mrX.print()

	mrX.updateName("Hii")
	mrX.print()
}

func (pPerson *person) updateName(name string) {
	(*pPerson).firstname = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
