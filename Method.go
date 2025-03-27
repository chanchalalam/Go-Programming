// Method example
package main

import "fmt"

// Define a struct
type person struct {
	name string
	age  int
}

// define a method with struct receiver
func (p person) display() {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)

}

func main() {
	//create instance of struc
	a := person{name: "a", age: 25}
	// call the method
	a.display()
}

package main

import "fmt"

type student struct {
	name   string
	stream string
}

func (s student) display() {
	fmt.Println("Name:", s.name)
	fmt.Println("Stream:", s.stream)
}

func main() {
	x := student{name: "Mirza Jasmine", stream: "BTech CSE"}
	x.display()
}

//Methods with Non-Struct Type Receiver

package main

import "fmt"

type number int

func (x number) square() number {
	return x * x
}
func main() {
	a := number(89)
	b := a.square()

	fmt.Println("Square of ", a, "is", b)
}

// Method with Pointer Reciever

package main

import "fmt"

type person struct {
	name string
}

func (p *person) changeName(newName string) {
	p.name = newName
}

func main() {
	a := person{name: "Chanchal"}
	fmt.Println("Before:", a.name)
	a.changeName("Jasmine")
	fmt.Println("After:", a.name)
}

// Methods Accepting Both Pointer and Value

package main

import "fmt"

type person struct {
	name string
}

func (p *person) updateName(newName string) {
	p.name = newName
}
func (p person) showName() {
	fmt.Println("Name:", p.name)
}

func main() {
	a := person{name: "Jasmine"}
	a.updateName("Jasmine")
	fmt.Println("After pointer method:", a.name)
	(&a).showName()
}


// create methods of the same name 
package main 
  
import "fmt"
  
// Creating structures 
type student struct { 
    name   string 
    branch string 
} 
  
type teacher struct { 
    language string 
    marks    int
} 
  
// Same name methods, but with 
// different type of receivers 
func (s student) show() { 
  
    fmt.Println("Name of the Student:", s.name) 
    fmt.Println("Branch: ", s.branch) 
} 
  
func (t teacher) show() { 
  
    fmt.Println("Language:", t.language) 
    fmt.Println("Student Marks: ", t.marks) 
} 
  
// Main function 
func main() { 
  
    // Initializing values 
    // of the structures 
    val1 := student{"Rohit", "EEE"} 
  
    val2 := teacher{"Java", 50} 
  
    // Calling the methods 
    val1.show() 
    val2.show() 
} 