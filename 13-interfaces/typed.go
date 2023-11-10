package main

import (
	"fmt"
	"math"
)

// Day-01
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

// Day-02
type Rectangle struct {
	Height  float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	return r.Breadth * r.Height
}

// Day-03
/*
func PrintArea(x interface{}) {
	switch shape := x.(type) {
	case Circle:
		fmt.Println("Area :", shape.Area())
	case Rectangle:
		fmt.Println("Area :", shape.Area())
	default:
		fmt.Println("Object has no area method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch shape := x.(type) {
	case interface{ Area() float32 }:  // any object that has an Area() method
		fmt.Println("Area :", shape.Area())
	default:
		fmt.Println("Object has no area method")
	}
}
*/

type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// Day-04 (introducing Perimeter)
func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Breadth + r.Height)
}

type PerimeterFinter interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinter) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// Day-05
/*
func PrintShape(x interface {
	AreaFinder
	PerimeterFinter
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinter
}

func PrintShape(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

// utility
func useCase(title string, fn func()) {
	fmt.Println(title)
	fn()
	fmt.Println("===========================")
}

func main() {
	// Day-01
	useCase("Day-1", func() {
		c := Circle{Radius: 12}
		fmt.Println("Area : ", c.Area())
	})

	// Day-02
	useCase("Day-2", func() {
		r := Rectangle{Height: 10, Breadth: 12}
		fmt.Println("Area : ", r.Area())
	})

	// Day-03
	useCase("Day-3", func() {
		c := Circle{Radius: 12}
		PrintArea(c)
		r := Rectangle{Height: 10, Breadth: 12}
		PrintArea(r)
	})

	// Day-04

	useCase("Day-4", func() {
		c := Circle{Radius: 12}
		PrintArea(c)
		PrintPerimeter(c)

		r := Rectangle{Height: 10, Breadth: 12}
		PrintArea(r)
		PrintPerimeter(r)
	})

	// Day-05 (encapsulate PrintArea() and PrintPerimeter())
	useCase("Day-5", func() {
		c := Circle{Radius: 12}
		PrintShape(c)

		r := Rectangle{Height: 10, Breadth: 12}
		PrintShape(r)
	})
}
