package calculator

import "fmt"

// private
/*
var opCount int

func OpCount() int {
	return opCount
}
*/

var opCount map[string]int

func init() {
	fmt.Println("init [calculator] - calc.go - 1")
	opCount = make(map[string]int)
}

func init() {
	fmt.Println("init [calculator] - calc.go - 2")
}

func OpCount() map[string]int {
	return opCount
}
