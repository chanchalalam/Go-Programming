// package main

// import "fmt"

// func main(){
// 	greeting := greet()

// 	fmt.Println(greeting)
// }

// func greet() string {
// 	return "Hello Chanchal! How R U ?"
// }


package main

import "testing"

func greet() string {
	return "Hello World"
}

func TestGreet(t *testing.T) {
	want := "Hello World"
	got := greet()
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}