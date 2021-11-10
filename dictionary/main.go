package main

import (
	"fmt"

	"github.com/tyomhk2015/go_basics/dictionary/dictionary"
)

func main() {
	myDictionary := dictionary.Dictionary{"One": "1"}

	fmt.Println(myDictionary.SearchWord("One"))
	fmt.Println(myDictionary.SearchWord("Two"))
	fmt.Println(myDictionary.Add("One", "9"))
	fmt.Println(myDictionary.Add("Hi", "Greetings"))
	fmt.Println(myDictionary.Add("Deckard", "Cain"))
	fmt.Println(myDictionary)
	myDictionary.Update("One", "Single")
	myDictionary.Update("Hi", "Bonjur")
	fmt.Println(myDictionary)
	myDictionary.Delete("Deckard")
	fmt.Println(myDictionary)
}