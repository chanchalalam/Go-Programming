// package main
// func ExampleMain(){
// 	main()
// }

package main

import "testing"

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

func TestGreetEnglish(t *testing.T) {
	lang := language("en") // A
	want := "Hello world"
	got := greet(lang) // B
	if got != want {   // C
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreetFrench(t *testing.T) {
	lang := language("fr") // A
	want := "Bonjour le monde"
	got := greet(lang) // B
	if got != want {   // C
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreetAkkadian(t *testing.T) {
	lang := language("akk") // A
	want := ""
	got := greet(lang) // B
	if got != want {   // C
		t.Errorf("expected: %q, got: %q", want, got)
	}
}