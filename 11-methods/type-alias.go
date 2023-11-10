package main

import "fmt"

// Create an alias for the "string" type
type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Laboris laborum id irure ullamco deserunt amet amet sint proident aliquip. Velit esse id ad est. Cillum pariatur minim sit reprehenderit enim anim labore cillum. Irure quis veniam laborum incididunt aute enim anim pariatur. Ullamco deserunt voluptate aliquip nulla magna consequat occaecat magna.")
	fmt.Println(str.Length())
}
