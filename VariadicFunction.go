// Variadic func in Go you to pass a variable number of arguments to a function.

// Product of Number
package main

import "fmt"

func multiply(num ...int) int {
	product := 1

	for _, n := range num {
		product *= n
	}
	return product
}

func main() {
	fmt.Println("product of 8,9,7:", multiply(8, 9, 7))
}

// Sum of Number

package main

import "fmt"

func sum(num ...int) int {
	sum := 0

	for _, n := range num {
		sum += n
	}
	return sum
}
func main() {
	fmt.Println("Sum of 9,9:", sum(9, 9))
}
