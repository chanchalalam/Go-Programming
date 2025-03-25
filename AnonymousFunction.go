// Anonymous Function is a function that does havea name.

package main

import "fmt"

func main() {
	// Anonymous function
	func() {
		fmt.Println("Me Cute cute")
	}()
}

//Assigning to a variable

package main

import "fmt"

func main() {
	value := func() {
		fmt.Println("ME CUTE CUTE")
	}
	value()
}

// Passing arguments in anonymous function

package main

import "fmt"

func main() {

	func(ele string) {
		fmt.Println(ele)
	}("Mirza Jasmine")
}

// Passing anonymous function as an argument

package main

import "fmt"

func MJ(i func(a, b string) string) {
	result := i("Jasmine", "Cute") // Call the passed function
	fmt.Println(result)            // Print the result
}

func main() {
	value := func(a, b string) string {
		return a + b + "Cute" // Concatenates a, b, and "Cute"
	}
	MJ(value) // Pass the function as an argument
}
