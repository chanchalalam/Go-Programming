package main

import "fmt"

func main(){
	greeting := greet("hn")

	fmt.Println(greeting)
}
type language string
func greet( l language) string {
	switch l {
	case "en":
		return "Hello World"
	case "fr":
		return "Bonjour le monde"
	case "hn":
		return "Aap sab ko Namaste"
	default:
		return ""
	}
}