package main

import "fmt"

func main() {
	var x int
	x = 100

	var intPtr *int
	intPtr = &x // value -> address (pointer)
	fmt.Println(intPtr)

	// dereferencing : address(pointer) -> value
	var y int
	y = *intPtr
	fmt.Println(y)

	fmt.Println(x == *(&x))

	var no int = 100
	fmt.Println("[@main] &no =", &no)
	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	var n1, n2 int = 100, 200
	fmt.Printf("before swapping: n1 = %d and n2 = %d\n", n1, n2)
	swap( /*  */ )
	fmt.Printf("after swapping: n1 = %d and n2 = %d\n", n1, n2)
}

func increment(n *int) {
	fmt.Println("[@ increment] &n =", n)
	*n += 1
}

func swap( /* ? */ ) /* no return result */ {
	/*  */
}
