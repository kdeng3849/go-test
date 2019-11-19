package main

import "fmt"

func adder() func(x int)int {
	sum := 0
	return func(x int) int {
		fmt.Println(sum,"+",x,"=")
		sum += x
		return sum
	}
}

func main() {
	sum := adder() // set to the anonymous function
	fmt.Printf("%T\n", sum)

	for i:=0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
