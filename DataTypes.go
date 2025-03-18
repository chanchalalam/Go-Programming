/*
Go Format Specifiers

%v	Default format	fmt.Printf("%v", 123)	123
%+v	Struct with field names	fmt.Printf("%+v", struct{Name string}{"Alice"})	{Name:Alice}
%#v	Go syntax representation	fmt.Printf("%#v", 123)	123
%T	Type of value	fmt.Printf("%T", 123)	int
%%	Literal percent sign	fmt.Printf("%%")	%
Boolean
%t	Boolean values	fmt.Printf("%t", true)	true
Integer
%b	Binary (base 2)	fmt.Printf("%b", 5)	101
%c	Unicode character	fmt.Printf("%c", 65)	A
%d	Decimal (base 10)	fmt.Printf("%d", 123)	123
%o	Octal (base 8)	fmt.Printf("%o", 8)	10
%O	Octal with 0o prefix	fmt.Printf("%O", 8)	0o10
%x	Hexadecimal (lowercase)	fmt.Printf("%x", 255)	ff
%X	Hexadecimal (uppercase)	fmt.Printf("%X", 255)	FF
%U	Unicode format	fmt.Printf("%U", 65)	U+0041
%q	Single-quoted character	fmt.Printf("%q", 65)	'A'
Floating-Point
%b	Scientific notation (powers of 2)	fmt.Printf("%b", 3.14)	4503599627370496p-50
%e	Scientific notation (lowercase)	fmt.Printf("%e", 123.456)	1.234560e+02
%E	Scientific notation (uppercase)	fmt.Printf("%E", 123.456)	1.234560E+02
%f	Decimal notation	fmt.Printf("%f", 123.456)	123.456000
%.nf	Fixed decimal places	fmt.Printf("%.2f", 123.456)	123.46
%g	Compact float (%e or %f)	fmt.Printf("%g", 123456.0)	123456
%G	Compact float (%E or %f)	fmt.Printf("%G", 123456.0)	123456
String & Slice
%s	String	fmt.Printf("%s", "Go")	Go
%q	Double-quoted string	fmt.Printf("%q", "Go")	"Go"
%x	Hex bytes (lowercase)	fmt.Printf("%x", "ABC")	414243
%X	Hex bytes (uppercase)	fmt.Printf("%X", "ABC")	414243
Pointer
%p	Pointer address	fmt.Printf("%p", &num)	0xc000014078

*/

// The use of integer

package main

import "fmt"

func main(){
	var X uint8 = 667
	fmt.Println(X,X-3)

	var Y int16 = 99999
	fmt.Println(Y+2,Y-2)
}

// Airthmetic Operation for Integers

package main

import "fmt"

func main() {

    var x int16 = 89
    var y int16 = 23
    //Addition
    fmt.Printf(" addition :  %d + %d = %d\n ", x, y, x+y)
    //Subtraction
    fmt.Printf("subtraction : %d - %d = %d\n", x, y, x-y)
    //Multiplication
    fmt.Printf(" multiplication : %d * %d = %d\n", x, y, x*y)
    //Division
    fmt.Printf(" division : %d / %d = %d\n", x, y, x/y)
    //Modulus
    fmt.Printf(" remainder : %d %% %d = %d\n", x, y, x%y)
}

//  The use of Floating point

package main

import "fmt"

func main (){
	a := 50.89
	b := 34.67

	c := a - b

	fmt.Printf("Result is : %f", c)
	fmt.Printf("\n The type of c: %T", c)
	fmt.Printf("\n The type of a: %T", a)
	fmt.Printf("\n The type of b: %T", b)
}

// Arithmetic operations for floating point numbers

package main

import "fmt"

func main() {
    var x float32 = 9.00
    var y float32 = 10.25
    //Addition
    fmt.Printf("addition :  %g + %g = %g\n ", x, y, x+y)
    //Subtraction
    fmt.Printf("subtraction : %g - %g = %g\n", x, y, x-y)
    //Multiplication
    fmt.Printf("multiplication : %g * %g = %g\n", x, y, x*y)
    //Division
    fmt.Printf("division : %g / %g = %g\n", x, y, x/y)

}

// The use of complex numbers

package main
import "fmt"

func main() {

   var a complex128 = complex(6, 2)
   var b complex64 = complex(9, 2)
   fmt.Println(a)
   fmt.Println(b)

   // Display the type
  fmt.Printf("The type of a is %T and "+
            "the type of b is %T", a, b)
}

// Build-in Func in complex Number

package main

import "fmt"

func main() {
	comp1 := complex(10, 11)
	comp2 := 13 + 33i

	fmt.Println("CN  of 1 is ", comp1)
	fmt.Println("CN of 2 is ", comp2)

	realNum := real(comp1)
	fmt.Println("RN OF CN 1 : ", realNum)
	imaginary := imag(comp2)
	fmt.Println("Imaginary of CN 2: ", imaginary)

}

// The use of Boolean

package main

import "fmt"

func main() {

	// variables
	str1 := "Mirza Jasmine"
	str2 := "mirzajasmine"
	str3 := "MirzaJasmine"
	result1 := str1 == str2
	result2 := str1 == str3

	// Display the result
	fmt.Println(result1)
	fmt.Println(result2)

	// Display the type of
	// result1 and result2
	fmt.Printf("The type of result1 is %T and "+
		"the type of result2 is %T",
		result1, result2)

}

// The use of Strings

package main

import "fmt"

func main() {
	str := "MirzaJasmine"

	fmt.Println("Length : ", len(str))
	fmt.Println("String : ", str)
	fmt.Printf("\nType of str is: %T", str)

	str1 := "Mirza"
	str2 := "Jasmine"

	fmt.Println("\nConcat:", str1+str2)
}
