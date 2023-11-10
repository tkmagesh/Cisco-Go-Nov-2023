// concurrent safe state management (using concurrent safe custom types)

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
		// increment(wg)
	}
	wg.Wait()
	fmt.Println("counter :", counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
