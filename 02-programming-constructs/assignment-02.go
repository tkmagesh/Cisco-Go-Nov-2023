/*
Build an interactive calculator

Display the following menu
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit

Accept the user choice

if the user choice == 5
	exit the application

if the user choice is 1 to 4
	accept 2 numbers
	perform the corresponding operation
	print the result
	IMPORTANT : display the menu again

if the user input is not (1-5)
	print "Invalid choice"
	IMPORTANT : display the menu again
*/

package main

import "fmt"

func main() {
	var userChoice, x, y, result int
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 4 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the operands:")
		fmt.Scanln(&x, &y)
		switch userChoice {
		case 1:
			result = x + y
		case 2:
			result = x - y
		case 3:
			result = x * y
		case 4:
			result = x / y
		}
		fmt.Println("Result :", result)
	}
}
