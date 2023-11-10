/* Inspect the resource utilization (using htop) */
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}

	flag.IntVar(&count, "count", 0, "Number of goroutines to schedule")
	flag.Parse()

	fmt.Printf("Starting %d goroutines... Hit ENTER to start!\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)    // increment the wg counter by 1
		go fn(wg, i) // scheduled to be executed in future
	}

	wg.Wait() //block the execution of this function until the wg counter becomes 0 (default is 0)
	fmt.Println("Main completed... Hit ENTER to shutdown!")
	fmt.Scanln()
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
