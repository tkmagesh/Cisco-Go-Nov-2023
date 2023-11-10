/*
communication between goroutines (using share memory by communicating [channels])
channels can also be used for synchronizing goroutines
*/

package main

import (
	"fmt"
	"time"
)

// consumer
func main() {

	fmt.Println("main started")
	ch := add(100, 200)
	result := <-ch
	fmt.Println("result :", result)
	fmt.Println("main completed")
}

// producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
