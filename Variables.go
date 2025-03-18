// Declaring a Variable
// In Go language variables are created in two different ways:

// Using var keyword: In Go language, variables are created using var keyword of a particular type, connected with name and provide its initial value.
// Syntax:

// var variable_name type = expression

// Example of Variable

package main

import "fmt"

func main() {
	var myvariable1 = 20
	var myvariable2 = " I m just cute little girl."
	var myvariable3 = 34.80
	var myvar1 int
	var myvar2 string
	var myvar3 float64
	var myva1, myva2, myva3 int = 2, 454, 6
	myvarb1 := 39
	myvarb2 := "ChanchalMirza"
	myvarb3 := 34.67
	myv1, myv2, myv3 := 800, 34, 56
	mv1, mv2 := 39, 45
	mv3, mv2 := 45, 100
	mr1, mr2, mr3 := 800, "Mirza", 47.56
	// Display the value and the
	// type of the variables
	fmt.Printf("The value of myvariable1 is : %d\n", myvariable1)
	fmt.Printf("The type of myvariable1 is : %T\n", myvariable1)
	fmt.Printf("The value of myvariable2 is : %s\n", myvariable2)
	fmt.Printf("The type of myvariable2 is : %T\n", myvariable2)
	fmt.Printf("The value of myvariable3 is : %f\n", myvariable3)

	fmt.Printf("The type of myvariable3 is : %T\n",
		myvariable3)

	fmt.Printf("The value of myvariable1 is : %d\n",
		myvar1)

	fmt.Printf("The value of myvariable2 is : %s\n",
		myvar2)

	fmt.Printf("The value of myvariable3 is : %f\n",
		myvar3)

	fmt.Printf("The value of myvariable1 is : %d\n",
		myva1)

	fmt.Printf("The value of myvariable2 is : %d\n",
		myva2)

	fmt.Printf("The value of myvariable3 is : %d\n",
		myva3)

	// Display the value and type of the variables
	fmt.Printf("The value of myvar1 is : %d\n", myvarb1)
	fmt.Printf("The type of myvar1 is : %T\n", myvarb1)
	fmt.Printf("The value of myvar2 is : %s\n", myvarb2)
	fmt.Printf("The type of myvar2 is : %T\n", myvarb2)
	fmt.Printf("The value of myvar3 is : %f\n", myvarb3)
	fmt.Printf("The type of myvar3 is : %T\n", myvarb3)
	fmt.Printf("The value of myvar1 is : %d\n", myv1)
	fmt.Printf("The type of myvar1 is : %T\n", myv1)
	fmt.Printf("The value of myvar2 is : %d\n", myv2)
	fmt.Printf("The type of myvar2 is : %T\n", myv2)
	fmt.Printf("The value of myvar3 is : %d\n", myv3)
	fmt.Printf("The type of myvar3 is : %T\n", myv3)
	fmt.Printf("The value of myvar1 and myvar2 is : %d %d\n", mv1, mv2)
	fmt.Printf("The value of myvar3 and myvar2 is : %d %d\n", mv3, mv2)
	fmt.Printf("The value of myvar1 is : %d\n", mr1)
	fmt.Printf("The type of myvar1 is : %T\n", mr1)
	fmt.Printf("The value of myvar2 is : %s\n", mr2)
	fmt.Printf("The type of myvar2 is : %T\n", mr2)
	fmt.Printf("The value of myvar3 is : %f\n", mr3)
	fmt.Printf("The type of myvar3 is : %T\n", mr3)
}
