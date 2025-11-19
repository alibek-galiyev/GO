package main

import "fmt"

func main() {
	// Declare and initialize with explicit type
	var coffeeName string = "Espresso"
	// Type inferred
	var size = "Small"
	var price float32 = 2.50
	// Short declaration and initialization. Possible only inside functions
	x := 10

	fmt.Println("Medium Espresso price is $2.50")
	fmt.Println(size, coffeeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)
	fmt.Println("x is", x)
}
