/* streaming data through channels */

package main

import (
	"fmt"
	"time"
)

func main() {
	var count int = 10
	ch := genFiboncci(count)
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done!")
}

func genFiboncci(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for x, y, i := 0, 1, 0; i < count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- x
			x, y = y, x+y
		}
	}()
	return ch
}
