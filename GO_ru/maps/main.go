package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":   2.20,
		"Americano":  3.00,
		"Latte":      4.00,
		"Cappuccino": 4.50,
		"Matcha":     4.00,
	}
	fmt.Println("Menu for today:", menu)
	fmt.Println("Latte costs:", menu["Latte"])

	menu["Latte"] += 1
	fmt.Println("New Latte price:", menu["Latte"])
	fmt.Println("Len:", len(menu))
	menu["Mocha"] = 6.25
	fmt.Println("Updated menu:", menu)

	drink := "Tea"
	fmt.Println(menu[drink])

	price, exists := menu[drink]
	if exists {
		fmt.Println("Price:", price)
	} else {
		fmt.Println(drink, "is not on menu!")
	}

	delete(menu, "Matcha")
	fmt.Println(menu)
	delete(menu, "Tea")

	for key, val := range menu {
		fmt.Printf("The price of %s is %.2f\n", key, val)
	}

	stock := make(map[string]int)

	if stock == nil {
		fmt.Printf("Stock map is nil")
	}

	stock["Latte"] = 23
	stock["Espresso"] = 100
	fmt.Println(stock)

	coffeeStock := map[string]int{
		"Espresso":   10,
		"Latte":      5,
		"Cappuccino": 8,
	}

	fmt.Println(coffeeStock)
	sellCoffee(coffeeStock, "Latte", 7)
	sellCoffee(coffeeStock, "Espresso", 3)
	sellCoffee(coffeeStock, "Cappuccino", 5)
	fmt.Println(coffeeStock)

}

func sellCoffee(stock map[string]int, coffeeType string, qty int) {
	currentStockQty, exists := stock[coffeeType]

	if exists {
		if currentStockQty >= qty {
			stock[coffeeType] -= qty
			fmt.Printf("Sold %d %s(s)\n", qty, coffeeType)
		} else {
			fmt.Printf("Not enough %s in stock. Current qty is %d\n", coffeeType, currentStockQty)
		}
	} else {
		fmt.Printf("%s is not available in stock!\n", coffeeType)
	}
}
