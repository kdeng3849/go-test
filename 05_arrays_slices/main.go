package main

import "fmt"

// arrays must be fixed size
// slices aren't fixed size

func main() {
	// arrays
	//var fruits [2]string
	//fruits[0] = "apple"
	//fruits[1] = "orange"

	//declare and assign simultaneously
	//var fruits = [2]string{"apple", "orange"}
	fruits := [2]string{"apple", "orange"}

	// slices
	fruitSlice := []string{"apple", "banana", "orange"}

	fmt.Println(fruits)          // [apple orange]
	fmt.Println(len(fruitSlice)) // 3
	fmt.Println(fruitSlice[1:2]) // [banana]
}
