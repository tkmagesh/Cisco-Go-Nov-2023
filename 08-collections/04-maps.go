package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["pen"] = 3
		productRanks["pencil"] = 1
		productRanks["marker"] = 5
	*/

	// productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	productRanks := map[string]int{
		"pen":    3,
		"pencil": 1,
		"marker": 5,
	}
	fmt.Println(productRanks)

	fmt.Println("len(productRanks) :", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Delete an item")
	keyToDelete := "notepad"
	delete(productRanks, keyToDelete) // delete operation is a "noop" if the key doesn't exists
	fmt.Println(productRanks)

	fmt.Println("check for the existence of a key")
	// keyToCheck := "pen"
	keyToCheck := "notepad"
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Value of key [%q] is %d\n", keyToCheck, val)
	} else {
		fmt.Printf("Key [%q] doesn't exists\n", keyToCheck)
	}
}
