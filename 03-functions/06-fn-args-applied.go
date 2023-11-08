package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		log.Println("invocation started")
		add(100, 200)
		log.Println("invocation completed")

		log.Println("invocation started")
		subtract(100, 200)
		log.Println("invocation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("multiply result :", x*y)
	}, 100, 200)

}

func logOperation(oper func(int, int), x, y int) {
	log.Println("invocation started")
	oper(x, y)
	log.Println("invocation completed")
}

func logAdd(x, y int) {
	log.Println("invocation started")
	add(x, y)
	log.Println("invocation completed")
}

func logSubtract(x, y int) {
	log.Println("invocation started")
	subtract(x, y)
	log.Println("invocation completed")
}

// 3rd party library
func add(x, y int) {
	fmt.Println("add result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("subtract result :", x-y)
}
