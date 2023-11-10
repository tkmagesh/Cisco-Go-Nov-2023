/* Using a sync.WaitGroup to synchronize goroutines */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1) // increment the wg counter by 1
		go f1(wg) // scheduled to be executed in future
	}
	f2()
	wg.Wait() //block the execution of this function until the wg counter becomes 0 (default is 0)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Println("f1 started")
	time.Sleep(7 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
