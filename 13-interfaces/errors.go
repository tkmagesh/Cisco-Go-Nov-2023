package main

import (
	"fmt"
)

// replace the following with a custom error type that has the multiplier & divisor data encapsulated
// var ErrDivideByZero error = errors.New("divisor cannot be 0")

type ErrDivideByZero struct {
	Multiplier int
	Divisor    int
}

// error interface implementation
func (e ErrDivideByZero) Error() string {
	return fmt.Sprintf("divide by zero error, multiplier = %d and divisor = %d", e.Multiplier, e.Divisor)
}

func NewErrDivideByZero(multiplier, divisor int) ErrDivideByZero {
	return ErrDivideByZero{
		Multiplier: multiplier,
		Divisor:    divisor,
	}
}

func main() {
	divisor := 0
	q, r, err := divide(100, divisor)
	if e, ok := err.(ErrDivideByZero); ok {
		fmt.Println("Do not attempt to divide by 0, err : ", e.Error())
		return
	}
	if err != nil {
		fmt.Println("something went wrong: ", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = NewErrDivideByZero(x, y)
		return
	}
	quotient, remainder = x/y, x%y
	return
}
