package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	cisco := &Organization{
		Name: "CISCO",
		City: "Bengaluru",
	}

	e1 := Employee{
		Id:   100,
		Name: "Magesh",
		Org:  cisco,
	}
	fmt.Printf("%+v\n", e1)

	e2 := Employee{
		Id:   200,
		Name: "Suresh",
		Org:  cisco,
	}
	fmt.Printf("%+v\n", e2)

	fmt.Println("After changing the cisco city to Mysuru")
	cisco.City = "Mysuru"
	fmt.Printf("%+v\n", cisco)
	fmt.Printf("%+v\n", e1.Org)
	fmt.Printf("%+v\n", e2.Org)
}
