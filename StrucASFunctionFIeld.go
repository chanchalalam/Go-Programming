// Struct with Method as a Function Field

package main

import "fmt"

type Person struct {
	Name  string
	Greet func() string
}

func main() {
	person := Person{
		Name: "Mirza",
	}
	// Assign the greet function after the person is defined
	person.Greet = func() string {
		return "Hello, " + person.Name
	}
	// Call the function field
	fmt.Println(person.Greet())
}

// Struct with Parameterized Function Field

package main
import "fmt"
type Person struct {
    Name  string
    Greet func(string) string
}
func main() {
    person := Person{
        Name: "mirza",
    }
    // Assign the greet function after the person is defined
    person.Greet = func(greeting string) string {
        return greeting + ", " + person.Name
    }
    // Call the function field with a parameter
    fmt.Println(person.Greet("Hi"))
}

// Struct with Multiple Function Fields
package main
import "fmt"
type Person struct {
    Name     string
    Greet    func(string) string
    Farewell func() string
}
func main() {
    person := Person{
        Name: "Chanchal",
    }
    // Assign the greet and farewell functions after the person is defined
    person.Greet = func(greeting string) string {
        return greeting + ", " + person.Name
    }
    person.Farewell = func() string {
        return "Goodbye, " + person.Name
    }
    // Call the function fields
    fmt.Println(person.Greet("Hello"))
    fmt.Println(person.Farewell())
}

