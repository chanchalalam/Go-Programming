// package main

// import "testing"

// func TestGreetEnglish(t *testing.T) {
// 	lang := language("en") // A
// 	want := "Hello world"
// 	got := greet(lang) // B
// 	if got != want {   // C
// 		// Mark this test as failed
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }

// func TestGreetFrench(t *testing.T) {
// 	lang := language("fr") // A
// 	want := "Bonjour le monde"
// 	got := greet(lang) // B
// 	if got != want {   // C
// 		// Mark this test as failed
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }

// func TestGreetAkkadian(t *testing.T) {
// 	// Akkadian is not implemented yet!
// 	lang := language("akk") // A
// 	want := ""
// 	got := greet(lang) // B
// 	if got != want {   // C
// 		// Mark this test as failed
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }



package main

import "fmt"

// Mock language function
func language(lang string) string {
	return lang
}

// Mock greet function
func greet(lang string) string {
	switch lang {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}

func main() {
	// Test Greet English
	lang := language("en")
	want := "Hello world"
	got := greet(lang)
	if got != want {
		fmt.Printf("Test Failed: expected %q, got %q\n", want, got)
	} else {
		fmt.Println("Test Passed: English")
	}

	// Test Greet French
	lang = language("fr")
	want = "Bonjour le monde"
	got = greet(lang)
	if got != want {
		fmt.Printf("Test Failed: expected %q, got %q\n", want, got)
	} else {
		fmt.Println("Test Passed: French")
	}

	// Test Greet Akkadian
	lang = language("akk")
	want = ""
	got = greet(lang)
	if got != want {
		fmt.Printf("Test Failed: expected %q, got %q\n", want, got)
	} else {
		fmt.Println("Test Passed: Akkadian (Not Implemented)")
	}
}