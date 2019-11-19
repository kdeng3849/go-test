package main

import "fmt"

func main() {
	// define map
	//emails := make(map[string]string)

	// assign key, values
	//emails["Kelly"] = "kelly@packet.com"
	//emails["Tammy"] = "tammy@packet.com"

	// assign key, values (shorthand)
	emails := map[string]string{"Kelly": "kelly@packet.com", "Tammy": "tammy@packet.com"}
	emails["hello"] = "hello@gmail.com"

	// delete from map
	delete(emails, "Kelly")

	fmt.Println(emails)
	fmt.Println(emails["Kelly"])
	fmt.Println(len(emails))
}