package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

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

	// appending new values
	nos = append(nos, 8)
	nos = append(nos, 7, 9)

	hundreds := []int{100, 200}
	nos = append(nos, hundreds...)
	fmt.Println("nos :", nos)

	dupNos := nos // creating a copy of the slice (both refer to the same instance)
	dupNos[0] = 9999
	fmt.Printf("nos[0] = %d, dupNos[0] = %d\n", nos[0], dupNos[0])

	sort(nos) // sort the nos array
	fmt.Println("After sorting, nos :", nos)

	// slicing
	fmt.Println("Slicing....")
	fmt.Println("nos[2:5] :", nos[2:5])
	fmt.Println("nos[2:] :", nos[2:])
	fmt.Println("nos[:5] :", nos[:5])
}

func sort(list []int) /* no return result */ {
	// use bubble sort to sort the given list of integers
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
