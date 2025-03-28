//  Anonymous structures are unnamed, temporary structures used for a one-time purpose, while anonymous fields allow embedding fields without names.

// AS Example

package main

import "fmt"

// Student struct with an anonymous inner structure for personal details
type Student struct {
	personalDetails struct { // Anonymous inner structure for personal details
		name       string
		enrollment int
	}
	GPA float64 // Standard field
}

func main() {
	// Initializing the anonymous structure inside Student
	student := Student{
		personalDetails: struct {
			name       string
			enrollment int
		}{
			name:       "A",
			enrollment: 12345,
		},
		GPA: 3.8,
	}
	// Displaying values
	fmt.Println("Name:", student.personalDetails.name)
	fmt.Println("Enrollment:", student.personalDetails.enrollment)
	fmt.Println("GPA:", student.GPA)
}

// Anonymous field

package main
import "fmt"

// Student struct using anonymous fields
type Student struct {
    int     // Enrollment number (anonymous field)
    string  // Name (anonymous field)
    float64 // GPA (anonymous field)
}

func main() {
    // Initializing the Student struct with anonymous fields
    student := Student{12345, "A", 3.8}

    // Displaying values
    fmt.Println("Enrollment:", student.int)
    fmt.Println("Name:", student.string)
    fmt.Println("GPA:", student.float64)
}

