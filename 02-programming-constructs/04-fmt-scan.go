package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter your name :")
		fmt.Scanln(&name)
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	var x, y int
	fmt.Println("Enter 2 numbers :")
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)

}
