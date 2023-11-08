package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)

		logMultiply := getLogOperation(func(i1, i2 int) {
			fmt.Println("multiply result :", i1*i2)
		})
		logMultiply(100, 200)
	*/

	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		profiledSubtract(100, 200)
	*/

	logAdd := getLogOperation(add)
	profiledLogAdd := getProfiledOperation(logAdd)
	profiledLogAdd(100, 200)
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("invocation started")
		oper(x, y)
		log.Println("invocation completed")
	}
}

func getProfiledOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Println("elapsed : ", elapsed)
	}
}

// 3rd party library
func add(x, y int) {
	fmt.Println("add result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("subtract result :", x-y)
}
