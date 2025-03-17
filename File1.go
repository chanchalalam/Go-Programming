// package main

// import (
// 	"fmt"
// 	"example.com/file2" // Importing the file2 package
// )

// // Exported variable (Accessible outside file1.go)
// var ExportedVariable = "Hello, World!"

// func main() {
// 	// Accessing the exported variable from file1.go
// 	fmt.Println(ExportedVariable)

// 	// Accessing the exported variable from file2.go
// 	fmt.Println(file2.AnotherExportedVariable)
// }

package main
 
import (
    "fmt"
)
 
// Exported variable
var ExportedVariable = "Hello, World!"
 
func main() {
    // Accessing exported identifier in the same file
    fmt.Println(ExportedVariable)
 
    // Accessing exported identifier from another package
    fmt.Println(file2.AnotherExportedVariable)
}