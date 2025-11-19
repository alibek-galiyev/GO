package main

import "fmt"

const discount float64 = 10

func main() {
	coffee := "Espresso"
	pointer := &coffee

	fmt.Println("Coffee name", coffee)
	fmt.Println("Memory location", &coffee)
	fmt.Printf("Pointer address: %p\n", pointer)

	coffeeCopy := coffee

	fmt.Println("Coffee name", coffeeCopy)
	fmt.Println("Memory location", &coffeeCopy)

	*pointer = "Cappuccino"
	fmt.Printf("Pointer address: %s\n", coffee)
	fmt.Printf("Pointer address: %s\n", coffeeCopy)

	coffeePrice := 5.00
	// discount := 10.

	coffeePrice = applyDiscount(coffeePrice)

	fmt.Printf("Coffee Price after discount: %.2f\n", coffeePrice)

	applyDiscountPointers(&coffeePrice)

	fmt.Printf("Coffee Price after discount: %.2f\n", coffeePrice)

}

func applyDiscount(price float64) float64 {
	return price * (1 - discount/100)
}

func applyDiscountPointers(price *float64) {
	*price = *price * (1 - discount/100.)
}
