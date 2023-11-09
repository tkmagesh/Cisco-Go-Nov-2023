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

	// accessing the members with a pointer
	/*
		var productPtr *Product = &product
		fmt.Println(productPtr.id)
	*/

	// use the applyDiscount function to update the cost of the product with 10% discount
	fmt.Println("After applying 10% discount")
	applyDiscount(&product, 10)
	display(product)
}

func display(p Product) {
	fmt.Printf("id = %d, name = %q, cost = %0.2f\n", p.id, p.name, p.cost)
}

func applyDiscount(productPtr *Product, discountPercentage float32) {
	productPtr.cost = productPtr.cost * ((100 - discountPercentage) / 100)
}
