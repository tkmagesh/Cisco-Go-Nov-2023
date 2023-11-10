// concurrent safe state management (using sync.Mutex)

package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

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

	// unsafe data manipulation
	// counter++

	// concurrent safe data manipulation
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()

}
