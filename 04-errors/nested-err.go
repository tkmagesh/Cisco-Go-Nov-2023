package main

import (
	"errors"
	"fmt"
)

var ErrE2 error = errors.New("error from f2")
var ErrE1 error = errors.New("error from f1")

func main() {
	err := f1()
	fmt.Println(err)
	if errors.Is(err, ErrE1) {
		fmt.Println("f1 returned error")
	}
	if errors.Is(err, ErrE2) {
		fmt.Println("f2 returned error")
	}
}

func f1() error {
	var e2Err = f2()
	// return errors.Join(ErrE1, e2Err)
	return fmt.Errorf("errors occurred : %w and %w", ErrE1, e2Err)
}

func f2() error {
	return ErrE2
	// return nil
}
