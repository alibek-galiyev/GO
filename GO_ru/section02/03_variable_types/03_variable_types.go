package main

import "fmt"

func main() {
	name := "Americano"
	price := 2.99
	ready := true
	count := 5

	fmt.Printf("Type of 'name' is a %T\n", name)
	fmt.Printf("Type of 'price' is a %T\n", price)
	fmt.Printf("Type of 'ready' is a %T\n", ready)
	fmt.Printf("Type of 'count' is a %T\n", count)
}
