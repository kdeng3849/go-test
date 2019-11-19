package main

import "fmt"

func main() {
	ids := []int{1,2,3,4,5}
	emails := map[string]string{"Kelly": "kelly@packet.com", "Tammy": "tammy@packet.com"}

	for i,id := range ids {
		fmt.Printf("%d: %d\n", i, id)
	}

	// not using index: put '_' as placeholder
	for _,id := range ids {
		fmt.Printf("%d\n", id)
	}

	for k,v := range emails {
		fmt.Printf("%s, %s\n", k, v)
	}

	for v := range emails {
		fmt.Printf("%s\n", v)
	}

}