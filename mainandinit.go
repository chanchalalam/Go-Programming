// The concept of main func

package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	s := []int{354, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice:", s)

	fmt.Println("Index value:", s)

	strings.Index("MirzaJasmine", "Jas")

	fmt.Println("Time: ", time.Now().Unix())
}

// the use of init

package main

import "fmt"

func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello! init() function")
}

func main() {
	fmt.Println("Welcome to main() function")
}
