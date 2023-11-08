package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(values ...int) int {
	result := 0
	for idx := 0; idx < len(values); idx++ {
		result += values[idx]
	}
	return result
}
