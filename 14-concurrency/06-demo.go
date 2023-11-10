/* communication between goroutines */

package main

import (
	"fmt"
	"sync"
	"time"
)

// communicate by sharing memory
var result int

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println("result :", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	result = x + y
}
