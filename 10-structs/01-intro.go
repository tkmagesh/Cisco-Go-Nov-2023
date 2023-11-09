package main

import "fmt"

func main() {
	/*
		var product struct {
			id   int
			name string
			cost float32
		}

		product.id = 100
		product.name = "Pen"
		product.cost = 10
		// fmt.Println(product)
		// fmt.Printf("%#v\n", product)
	*/

	var product struct {
		id   int
		name string
		cost float32
	} = struct {
		id   int
		name string
		cost float32
	}{
		id:   100,
		name: "pen",
		cost: 10,
	}

	// fmt.Printf("%+v\n", product)
	display(product)
}

func display(p struct {
	id   int
	name string
	cost float32
}) {
	fmt.Printf("id = %d, name = %q, cost = %0.2f\n", p.id, p.name, p.cost)
}
