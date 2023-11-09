package calculator

import "fmt"

func init() {
	fmt.Println("init [calculator] - add.go")
}

// public
func Add(x, y int) int {
	// opCount++
	opCount["Add"]++
	return x + y
}
