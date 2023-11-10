/* communication between goroutines (using share memory by communicating [channels]) */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch
	fmt.Println("result :", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	result := x + y
	ch <- result // send the result through the channel
}
