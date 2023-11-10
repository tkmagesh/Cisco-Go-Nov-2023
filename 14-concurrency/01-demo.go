package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduled to be executed in future
	f2()
	// DO NOT DO THIS (coz this is a very primitive way of synchronizing goroutines)
	time.Sleep(2 * time.Second) //blocking the execution of this function so that the scheduler can look for other goroutines scheduled and execute them
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
