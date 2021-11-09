package main

// fmt: format
import (
	"fmt"
	"strings"
)

func main() {
	// Main & Import
	fmt.Println("Hello Go!")

	// Var
	// var age int = 22 (Long declaration)
	age := 22 // (Shorthand declaration)
	age++

	// Const (Cannot use shorthand)
	// https://medium.com/technofunnel/creating-constant-in-golang-c958c1821c20
	const name string = "Louise"
	// name = "Deckard Cain"

	fmt.Println(name, age, multiply(3,3))

	//A Function with multiple value return.
	length, upperName, lowerName := caseAndUpper(name)

	// '_' : ignored values, compiler will ignore this.
	// Useful for getting a few of returned values.
	_, bigName, smallName := caseAndUpper(name)
	fmt.Println(length, upperName, lowerName)
	fmt.Println(bigName, smallName)

	// A function with many arguments
	repeatWords("Sushi", "Hamburger", "Yakitori", "Gellato", "Dalgona", "Squid", "Takoyaki")

	// A Naked function
	menu, bill := orderUberEats("Spicy Pizza")
	fmt.Println(menu, bill)

	// A Defer function
	alarmUberEats()

	// For loop w/ range
	fmt.Println(accumulate(1,2,3,4,5))

	// If & Variable Expression
	fmt.Println(areYouABoomer(20), areYouABoomer(26))

	// Switch
	fmt.Println(sushiGrade(1500))
	fmt.Println(sushiGrade(29000))

	// Pointer
	address()

	// Array * Slices
	nexonGames()

	// Maps
	capitalCities()
}

// The last 'int', at the right, is the return 'type' of this function.
func multiply(a int, b int) int {
	return a * b
}

// A Function with multiple value return.
func caseAndUpper(name string) (int, string, string) {
	return len(name), strings.ToUpper(name), strings.ToLower(name)
}

// A function with many arguments, just like spread operator in JS.
func repeatWords(words ...string) {
	fmt.Println(words)
}

// A naked function.
// Return variables are declared at the beginning.
func orderUberEats(menu string) (food string, payment int) {
	food = menu
	payment = len(menu) * 1000
	return
}

// A 'defer' function.
// A function excecutes after the previous function finishes its job.
func alarmUberEats() {
	defer fmt.Println("The delivery has arrived at the destination.");
	fmt.Println("The delivery is on the way.");
}

// For loop w/ range.
// The 1st return of range is the index of the array.
// The 2nd return of range is the value of the array.
func accumulate(numbers ...int) int {
	total := 0
	for index, value := range numbers {
		total += value
		fmt.Println(index, value)
	}
	return total
}

// If
func areYouABoomer(age int) bool {
	// Normal If
	if age > 30 {
		return true
	}

	// Variable Expression
	if newAge := age + 5; newAge > 30 {
		fmt.Println("Variable Expression")
		return true
	}
	return false
}

// Switch
func sushiGrade(price int) string {
	// Normal switch
	switch price {
		case 700:
			return "Supermarket sushi"
		case 800:
			return "Supermarket sushi with extra sashimi"
	}

	// Switch with flexibility
	switch{
		case price > 800 && price < 2500:
			return "Restaurant sushi"
		case price >= 2500 && price < 5000:
			return "Superior restaurant sushi"
	}

	// Switch with variable expression
	switch newPrice := price + 1000; newPrice {
		case 10000:
			return "Omakase sushi"
		case 30000:
			return "Premium omakase sushi"
	}
	return "Enter different price."
}

// Pointer
func address() {
	// &: Get the memory address of the variable
	// *: See what is inside the memory address
	chiba := 272
	otaku := 145
	tokyo := 100
	destination := &chiba
	fmt.Println(destination, *destination)
	chiba = otaku
	fmt.Println(destination, *destination)
	destination = &tokyo
	fmt.Println(destination, *destination, chiba)
}

// Array & Slices
func nexonGames() {
	// Normal array
	games := [4] string {"Mabinogi", "Vindictus", "MapleStory", "Arad Chronicle"}

	// Slice & appens
	consoleGames := [] string {"Persona 4G", "DragonQuest XI S", "Biohazard Village"}
	consoleGames = append(consoleGames, "Yakuza 0")
	fmt.Println(games, consoleGames)

	for index, value := range consoleGames {
		fmt.Println(index, value)
	}
}

// Maps
func capitalCities() {
	cities := map[string]string {"South Korea": "Seoul", "Japan": "Tokyo", "Thailand": "Bangkok", "Austria": "Wien"}

	// The order of newly added key-value is not guaranteed.
	cities["France"] = "Paris"

	// Iterate maps
	for _, value := range cities {
		fmt.Println(value)
	}
}