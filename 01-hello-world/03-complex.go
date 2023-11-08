package main

import "fmt"

func main() {
	var c1 complex128
	c1 = 4 + 5i
	fmt.Println("c1 :", c1)
	fmt.Println("real(c1) :", real(c1))
	fmt.Println("imag(c1) :", imag(c1))

	// complex type arithmatic
	var c2 complex128 = 7 + 9i
	c3 := c1 + c2
	fmt.Println(c3)
}
