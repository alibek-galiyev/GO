package main

import "fmt"

func main() {
	// Explicit type declaration
	var cupsQty int = 3
	// cupsQty = "four" // Compile-time error
	fmt.Println("Cups Qty:", cupsQty)

	// Implicit type declaration
	var wasProcessed = true
	fmt.Println("Order was processed:", wasProcessed)

	price := 4.50

	qty := 15

	total := price * float64(qty)

	fmt.Printf("Total earnings: %.2f\n", total)

	var coffee, milk, sugar = 1, 2, 1
	fmt.Println(coffee, milk, sugar)

	const pi = 3.14

	fmt.Printf("Area of the circle with radius 5 is: %.2f\n", pi*5*5)
	fmt.Printf("Type of pi is %T\n", pi)

	var (
		customerName string = "Alibek"
		tableNumber  int    = 8
		isReadyToPay bool   = false
	)

	const (
		Pi = 3.1415
		C  = 30000000000
	)
	fmt.Println(customerName, tableNumber, isReadyToPay)
	fmt.Println(Pi, C)
}
