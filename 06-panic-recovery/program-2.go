/* convert the panic (from divide) into an error (in the divideClient) so that that application can take a different course of action */
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by 0 error")

func main() {
	defer func() {
		fmt.Println("	[deferred @ main]")
		if err := recover(); err == nil {
			fmt.Println("App executed successfully")
			return
		} else {
			fmt.Println("panic occurred :", err)
			return
		}

	}()
	divisor := 0
	fmt.Println("[@main] initiating divide")
	q, r, err := divideClient(100, divisor)
	if err == nil {
		fmt.Printf("[@main] Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else {
		fmt.Println("try again! err :", err)
	}
}

// intermediate function to conver the panic into an error
func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party (a possibility of a panic)
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(errors.New("divide by 0 error"))
	}
	quotient, remainder = x/y, x%y
	return
}
