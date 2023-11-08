package main

import "fmt"

// package level unused variable
var unused_var int

func main() {
	/*
		var name string
		name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	/*
		var name string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	// type inference
	/*
		var name = "Magesh" // the variable is determined by the data type being assigned
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	// using ":="
	name := "Magesh" // the variable is determined by the data type being assigned
	fmt.Printf("Hi %s, Have a nice day!\n", name)

	// unused variable
	// n := 100
	var n int
	n = 100
	fmt.Println(n) //reading the value of the variable makes it being "used"

	//multiple variables
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// declaration & initialization
	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// type inference
	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)
}
