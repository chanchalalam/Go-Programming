package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) display() {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)

}

func main() {
	a := person{name: "a", age: 25}

	a.display()
}
