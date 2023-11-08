package main

import "fmt"

func main() {
	// const pi float32 = 3.14
	const pi = 3.14
	fmt.Println("pi :", pi)

	// combining multiple complex declarations
	// note: constants can remain unused
	const (
		x = 100
		y = 200
	)
}
