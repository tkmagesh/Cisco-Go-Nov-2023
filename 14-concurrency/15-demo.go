/* streaming data through channels */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genFiboncci()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		break
	}
	fmt.Println("Done!")
}

func genFiboncci() <-chan int {
	count := rand.Intn(20)
	ch := make(chan int)
	go func() {
		for x, y, i := 0, 1, 0; i < count; i++ {
			ch <- x
			x, y = y, x+y
		}
		fmt.Println("All the values produced")
		close(ch)
	}()
	return ch
}
