package main

import (
	"fmt"

	// "github.com/tkmagesh/cisco-go-nov-2023/09-modularity/calculator"
	"github.com/fatih/color"
	calc "github.com/tkmagesh/cisco-go-nov-2023/09-modularity/calculator"
	"github.com/tkmagesh/cisco-go-nov-2023/09-modularity/calculator/utils"
)

func main() {
	// fmt.Println("modularity demo")
	color.Red("modularity demo")
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("OpCount :", calculator.OpCount())
	*/

	// using package alias
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("OpCount :", calc.OpCount())
	fmt.Println(utils.GetCurrentTime())

}
