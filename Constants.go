//  How to declare: Constants are declared like variables but in using a const keyword as a prefix to declare a constant with a specific type.
//  It cannot be declared using “:=” syntax.

package main

import "fmt"

const PI = "Cute"

func main() {
	const NAME = "ChanchalAlam"
	fmt.Println("Hello", NAME)

	fmt.Println("Happy", PI, "Day")
	const Correct = true

	fmt.Println("Be Happy Always", Correct)
}


package main
 
import "fmt"
 
func main()
{
    const A = "GFG"
    var B = "GeeksforGeeks"
     
    // Concat strings.
    var helloWorld = A+ " " + B
    helloWorld += "!"
    fmt.Println(helloWorld) 
     
    // Compare strings.
    fmt.Println(A == "GFG")   
    fmt.Println(B < A) 
}


package main
 
import "fmt"
 
const Pi = 3.14
 
func main() 
{
    const trueConst = true
     
    // Type definition using type keyword
    type myBool bool    
    var defaultBool = trueConst // allowed
    var customBool myBool = trueConst // allowed
     
    //  defaultBool = customBool // not allowed
    fmt.Println(defaultBool)
    fmt.Println(customBool)   
}