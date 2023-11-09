package main

import "fmt"

func main() {
	// nos := []int{}
	// nos := make([]int, 0, 3)
	nos := make([]int, 2, 3)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 10)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 20)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 30)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 40)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 50)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)
}
