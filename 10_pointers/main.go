package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println("a: ",a,", b: ",b)
	fmt.Printf("type of b: %T\n", b)

	*b = 10
	fmt.Println("a: ", a)
}