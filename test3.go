package main

import (" fmt ")

func main(){
	greeting := greet("en")

	fmt.Println(greeting)

}

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε", // Greek
	"en": "Hello world", // English
	"fr": "Bonjour le monde", // French
	"שלום עולם" :"he", // Hebrew
	"vi": "Xin chào Thế Giới", // Vietnamese
	}

	func greet( l language) string{
		greeting, ok:= phrasebook[l]
		if !ok{
			return fmt.Sprintf("unsupported language: %q", l)

		}
		return greeting
	}