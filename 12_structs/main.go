package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	//first string
	//last string
	//age int
	//gender string

	first, last, gender string
	age int
}

// value receiver (just reading values)
func (p Person) greet() string {
	return "Hello, my name is " + p.first + " " + p.last + " and I'm " + strconv.Itoa(p.age)
}

// pointer receiver (changing values)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.last = spouseLastName
	}
}

func main()  {
	// long way (like a map)
	person := Person{first: "Kelly", last: "Deng", age: 21, gender: "F"}

	// short way
	person2 := Person{"Tammy", "Chan", "F", 21}

	fmt.Println(person)
	fmt.Println(person.first)
	person.first = "Avacadro"
	fmt.Println(person.first)
	fmt.Println(person2.last)

	fmt.Println(person.greet())
	person.hasBirthday()
	person.getMarried("Crunk")
	fmt.Println(person.greet())
}
