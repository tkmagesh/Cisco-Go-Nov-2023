// concurrent safe state management (using sync/atomic package)

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
		// increment(wg)
	}
	wg.Wait()
	fmt.Println("counter :", counter)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1)
}
