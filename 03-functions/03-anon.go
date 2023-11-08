package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	msg := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a good day!", userName)
	}("Suresh")
	fmt.Println(msg)

	addResult := func(x int, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("addResult :", addResult)

	q, r := func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
