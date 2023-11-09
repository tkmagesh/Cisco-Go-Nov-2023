package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	[deferred @ main]")
		if err := recover(); err == nil {
			fmt.Println("App executed successfully")
			return
		} else {
			fmt.Println("panic occurred :", err)
			return
		}

	}()
	divisor := 0
	fmt.Println("[@main] initiating divide")
	q, r := divide(100, divisor)
	fmt.Printf("[@main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) {
	defer func() {
		fmt.Println("	[deferred @ divide]")
	}()
	fmt.Println("[@divide] calculating quotient")
	quotient = x / y
	fmt.Println("[@divide] calculating remainder")
	remainder = x % y
	fmt.Println("[@divide] return the result")
	return
}
