package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong: ", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

// returning a generic error object
/*
func divide(x, y int) (quotient, remainder int , err error ) {
	if y == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient, remainder = x/y, x%y
	return
}
