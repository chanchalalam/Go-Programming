/// Struc in another struc

package main

import (
	"fmt"
)

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
}

func main() {
	p := Person{
		FirstName: "Mirza",
		LastName:  "Jasmine",
		Age:       18,
		Address: Address{
			Street:     "MDB",
			City:       "MDB",
			State:      "Bihar",
			PostalCode: "123456",
		},
	}

	fmt.Println(p.FirstName, p.LastName)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:")
	fmt.Println("Street:", p.Address.Street)
	fmt.Println("City:", p.Address.City)
	fmt.Println("State:", p.Address.State)
	fmt.Println("Postal Code:", p.Address.PostalCode)
}
