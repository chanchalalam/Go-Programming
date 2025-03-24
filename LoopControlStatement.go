// Loop Control Statements in Go Language

// Break Statement

package main

import "fmt"

// Main function
func main() {
	for i := 0; i < 5; i++ {

		fmt.Println(i)
		if i == 3 {
			break
		}
	}

}


// Goto Statement

// the use of goto statement 
package main 
  
import "fmt"
  
func main() { 
   var x int = 0 
     
   // for loop work as a while loop 
  Lable1: for x < 8 { 
      if x == 5 { 
           
         // using goto statement 
         x = x + 1; 
         goto Lable1 
      } 
      fmt.Printf("value is: %d\n", x); 
      x++;      
   }   
} 


// Continue Statement

package main 
  
import "fmt"
  
func main() { 
   var x int = 0 
     
   // for loop work as a while loop 
   for x < 8 { 
      if x == 5 { 
           
         // skip two iterations 
         x = x + 2; 
         continue; 
      } 
      fmt.Printf("value is: %d\n", x); 
      x++;      
   }   
} package main 
  
import "fmt"
  
func main() { 
   var x int = 0 
     
   // for loop work as a while loop 
   for x < 8 { 
      if x == 5 { 
           
         // skip two iterations 
         x = x + 2; 
         continue; 
      } 
      fmt.Printf("value is: %d\n", x); 
      x++;      
   }   
} 

