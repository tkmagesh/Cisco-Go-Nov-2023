package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Anim laborum qui ipsum dolor duis do ex."
	x = 99.99
	x = true
	x = struct{}{}
	fmt.Println(x)

	x = 100
	// y := x + 200
	// x = "Voluptate nisi excepteur quis tempor nulla nisi dolor minim tempor adipisicing anim qui voluptate nostrud."
	// y := x.(int) + 200

	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// x = 100
	// x = 99.99
	// x = true
	// x = "Sit in dolore id tempor aliqua aute magna."
	// x = struct{}{}
	x = 10 + 15i
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case float64:
		fmt.Println("x is a float64, x * 0.99 =", val*0.99)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case struct{}:
		fmt.Println("x is a struct{}")
	default:
		fmt.Println("x is an unknown type")
	}

}
