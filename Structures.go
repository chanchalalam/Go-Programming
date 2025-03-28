// A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type.

/* Declare the struc

type Address struct{

name string
street string
city string
state sttring
age int

}


type Address struct {
    name, street, city, state string
    age int
}
*/

// To Define a structure

//  Syntax - var a Address

package main

import "fmt"

type Address struct {
	name string
	city string
	age  int
}

func main() {
	var a Address
	fmt.Println(a)

	a1 := Address{name: "Chanchal", city: "Gurgaon", age: 18}
	fmt.Println("Address1:", a1)

	a2 := Address{name: "Kartik", city: "Gurgaon", age: 21}
	fmt.Println("Address2:", a2)
}

/* How to access fields of a struct?

To access individual fields of a struct you have to use dot (.) operator. */

package main

import "fmt"

// defining the struct
type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

// Main Function
func main() {
	c := Car{Name: "Nano", Model: "GTC4",
		Color: "Pink", WeightInKg: 1920}
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)
	c.Color = "Pink"
	fmt.Println("Car: ", c)
}

//Pointers to a struct

package main

import "fmt"

type Employee struct {
	firstName, lastName, role string
	age, salary               int
}

func main() {
	emp1 := &Employee{"Chanchal", "Alam", "Associate Data Scientist", 18, 25000}

	fmt.Println("First Name:", (*emp1).firstName)
	fmt.Println("Last Name:", (*emp1).lastName)
	fmt.Println("Role:", (*emp1).role)
	fmt.Println("Age:", (*emp1).age)
	fmt.Println("Salary:", (*emp1).salary)

}

// Structure Equality in Golang
// concept of struct equality
// using == operator

package main

import "fmt"

// Creating a structure
type Author struct {
	name      string
	branch    string
	language  string
	Practices int
}

// Main function
func main() {

	// Creating variables
	// of Author structure
	a1 := Author{
		name:      "Mirza",
		branch:    "CSE",
		language:  "Python",
		Practices: 38,
	}

	a2 := Author{
		name:      "Mirza",
		branch:    "CSE",
		language:  "Python",
		Practices: 38,
	}

	a3 := Author{
		name:      "Jasmine",
		branch:    "CSE",
		language:  "Python",
		Practices: 38,
	}

	if a1 == a2 {

		fmt.Println("Variable a1 is equal to variable a2")

	} else {

		fmt.Println("Variable a1 is not equal to variable a2")
	}

	if a2 == a3 {

		fmt.Println("Variable a2 is equal to variable a3")

	} else {

		fmt.Println("Variable a2 is not equal to variable a3")
	}
}

// Using DeepEqual() method

package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	name      string
	branch    string
	language  string
	Practices int
}

func main() {
	a1 := Author{
		name:      "Mirza",
		branch:    "CSE",
		language:  "Python",
		Practices: 38,
	}
	a2 := Author{
		name:      "Mirza",
		branch:    "CSE",
		language:  "Python",
		Practices: 38,
	}
	a3 := Author{
		name:      "Jasmine",
		branch:    "CSE",
		language:  "Python",
		Practices: 38,
	}
	fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a2))
	fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a3))
}
