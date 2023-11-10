/* channel behaviors */

package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)

	/*
		ch := make(chan int)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			data := <-ch
			fmt.Println(data)
			wg.Done()
		}()
		ch <- 100
		wg.Wait()
	*/
}

// modify the above to perform the "receive" operation in a goroutine
