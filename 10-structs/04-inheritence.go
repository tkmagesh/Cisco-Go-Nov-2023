package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	// Id string //overriding the Product.Id
	Product
	// Dummy
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}
func main() {
	// grapes := PerishableProduct{Product{100, "Grapes", 400}, "2 Days"}
	grapes := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Grapes",
			Cost: 400,
		},
		/*
			Dummy: Dummy{
				Id: 9999,
			},
		*/
		Expiry: "2 Days",
		// Id:     "Fruits-101",
	}

	// accessing the attributes of the base type
	// fmt.Println(grapes.Product.Id)
	fmt.Println(grapes.Product.Id)
	// fmt.Println(grapes.Dummy.Id)

	milk := NewPerishableProduct(102, "Milk", 40, "1 Day")
	fmt.Println(milk)
}
