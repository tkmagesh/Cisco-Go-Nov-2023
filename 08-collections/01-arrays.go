package main

import "fmt"

func main() {
	// var nos [5]int // memory allocated and initialized with the default value of int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// use :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println("nos :", nos)

	// len()
	fmt.Println("len(nos) :", len(nos))

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for idx := 0; idx < 5; idx++ {
		fmt.Printf("nos[%d] = %d, %v\n", idx, nos[idx], &nos[idx])
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	dupNos := nos // creating a copy of the array
	dupNos[0] = 9999
	fmt.Printf("nos[0] = %d, dupNos[0] = %d\n", nos[0], dupNos[0])

	sort( /* ? */ ) // sort the nos array
	fmt.Println("After sorting, nos :", nos)
}

func sort( /* ? */ ) /* no return result */ {
	// use bubble sort to sort the given list of integers
}
