package main

import "fmt"

// shorthand won't work outside of functions
// age := 21

func main() {
	var first, last = "Kelly", "Deng"
	age := 21.5

	fmt.Println(first, last, age)
	fmt.Printf("%T\n", age)
}