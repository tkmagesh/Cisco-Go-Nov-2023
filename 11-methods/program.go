package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (productPtr *Product) ApplyDiscount(discountPercentage float32) {
	productPtr.Cost = productPtr.Cost * ((100 - discountPercentage) / 100)
}

// inheriting methods
type PerishableProduct struct {
	Product
	Expiry string
}

// overriding methods
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {
	pen := Product{Id: 100, Name: "Pen", Cost: 10}
	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	// perishable product
	grapes := PerishableProduct{
		Product: Product{
			Id:   200,
			Name: "Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	fmt.Println(grapes.Format())
	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}
