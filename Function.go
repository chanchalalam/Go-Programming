// In Go, functions are blocks of code that perform specific tasks, which can be reused throughout the program to save memory,
//  improve readability, and save time. Functions may or may not return a value to the caller.

// Syntax

// func function_name(Parameter-list)(Return_type) {
//     // function body...
// }

// package main

// import "fmt"

// func multiply(a, b int) int {
// 	return a * b
// }

// func main() {
// 	result := multiply(5, 10)
// 	fmt.Printf("multiplication: %d", result)

// }

// Func - keyword to declare a func.
// func_name - the name of func .
// Parameter-list - a, b int parameter with their types
// return-type - int specifies the return type

// Function Arguments

// Call by Value

// package main

// import "fmt"

// func multiply(a, b int) int {
// 	a = a * 2
// 	return a * b
// }

// func main() {
// 	x := 10
// 	y := 55

// 	fmt.Printf("Before: x = %d, y = %d\n", x, y)
// 	result := multiply(x, y)
// 	fmt.Printf("multiplication : %d\n", result)
// 	fmt.Printf("After: x = %d , y = %d\n", x, y)
// }

// Call by Reference

package main

import "fmt"

func multiply(a, b *int) int {
	*a = *a * 2
	return *a * *b
}

func main() {
	x := 5
	y := 10
	fmt.Printf("Before: x = %d, y = %d\n", x, y)
	result := multiply(&x, &y)
	fmt.Printf("multiplication: %d\n", result)
	fmt.Printf("After: x = %d, y = %d\n", x, y)
}
