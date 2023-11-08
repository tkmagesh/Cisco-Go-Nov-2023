package main

import "fmt"

func main() {
	fmt.Println("for[v1]")
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("for[v2] (while)")
	no := 1
	for no < 100 {
		no += no
	}
	fmt.Println("no :", no)
}
