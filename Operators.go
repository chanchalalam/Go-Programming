package main

import "fmt"

func main() {
	p := 55
	q := 20
	//Arithmetic Operators

	// Addition
	result1 := p + q
	fmt.Printf("Result of p + q = %d", result1)

	// Subtraction
	result2 := p - q
	fmt.Printf("\nResult of p - q = %d", result2)

	// Multiplication
	result3 := p * q
	fmt.Printf("\nResult of p * q = %d", result3)

	// Division
	result4 := p / q
	fmt.Printf("\nResult of p / q = %d", result4)

	// Modulus
	result5 := p % q
	fmt.Printf("\nResult of p %% q = %d", result5)

	// Relational Operators
	// ‘=='(Equal To)
	result6 := p == q
	fmt.Println(result6)

	// ‘!='(Not Equal To)
	result7 := p != q
	fmt.Println(result7)

	// ‘<‘(Less Than)
	result8 := p < q
	fmt.Println(result8)

	// ‘>'(Greater Than)
	result9 := p > q
	fmt.Println(result9)

	// ‘>='(Greater Than Equal To)
	result10 := p >= q
	fmt.Println(result10)

	// ‘<='(Less Than Equal To)
	result11 := p <= q
	fmt.Println(result11)

	//Logical Operators

	if p != q && p <= q {
		fmt.Println("True")
	}

	if p != q || p <= q {
		fmt.Println("True")
	}

	if !(p == q) {
		fmt.Println("True")
	}

	//Bitwise Operators

	// & (bitwise AND)
	result011 := p & q
	fmt.Printf("Result of p & q = %d", result011)

	// | (bitwise OR)
	result12 := p | q
	fmt.Printf("\nResult of p | q = %d", result12)

	// ^ (bitwise XOR)
	result13 := p ^ q
	fmt.Printf("\nResult of p ^ q = %d", result13)

	// << (left shift)
	result14 := p << 1
	fmt.Printf("\nResult of p << 1 = %d", result14)

	// >> (right shift)
	result15 := p >> 1
	fmt.Printf("\nResult of p >> 1 = %d", result15)

	// &^ (AND NOT)
	result16 := p &^ q
	fmt.Printf("\nResult of p &^ q = %d", result16)

	//Assignment Operator

	// “=”(Simple Assignment)
	p = q
	fmt.Println(p)

	// “+=”(Add Assignment)
	p += q
	fmt.Println(p)

	//“-=”(Subtract Assignment)
	p -= q
	fmt.Println(p)

	// “*=”(Multiply Assignment)
	p *= q
	fmt.Println(p)

	// “/=”(Division Assignment)
	p /= q
	fmt.Println(p)

	// “%=”(Modulus Assignment)
	p %= q
	fmt.Println(p)

	//Misc Operator

	// 	   & - return address of variable
	// 	   * - gives pointer to a variable
	// 	   <- - name of this operator is receive. It is used to receive a value from the channel.
	//

	a := 4
	// Using address of operator(&) and
	// pointer indirection(*) operator
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
