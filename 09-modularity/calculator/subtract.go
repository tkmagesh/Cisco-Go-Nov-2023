package calculator

import "fmt"

func init() {
	fmt.Println("init [calculator] - subtract.go")
}

// public
func Subtract(x, y int) int {
	// opCount++
	opCount["Subtract"]++
	return x - y
}
