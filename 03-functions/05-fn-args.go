package main

import "fmt"

func main() {
	exec_fn(f1)
	exec_fn(f2)
}

func exec_fn(fn func()) {
	// execute f1 OR f2 based on the input
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
