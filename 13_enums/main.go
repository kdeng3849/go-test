package main

import "fmt"

type Weekday int

const(
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (s Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[s]
}

func main() {
	fmt.Printf("I work %ss/%ss.",Tuesday, Wednesday)
}
