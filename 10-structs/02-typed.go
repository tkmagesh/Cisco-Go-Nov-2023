package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
}

func main() {
	/*
		var product Product = Product{
			id:   100,
			name: "pen",
			cost: 10,
		}
	*/

	// partially initializing the instance
	/*
		var product Product = Product{
			id:   100,
			name: "pen",
		}
	*/

	// CANNOT partially initialize using the following syntax
	// var product Product = Product{100, "Pen", 10}

	product := Product{
		id:   100,
		name: "pen",
		cost: 10,
	}
	display(product)
}

func display(p Product) {
	fmt.Printf("id = %d, name = %q, cost = %0.2f\n", p.id, p.name, p.cost)
}
