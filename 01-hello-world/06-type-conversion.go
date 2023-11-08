package main

import "fmt"

func main() {
	var i int32 = 100
	var f float64 = 200.99

	// syntax for type conversion : use the typename as a function
	f = f + float64(i)
	fmt.Println(f)

}
