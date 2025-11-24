package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType   string
	Size         string
	CustomerName string
	BonusPoints  int
}

func main() {
	var order1 CoffeeOrder
	order1.CoffeeType = "Latte"
	order1.Size = "Large"
	order1.CustomerName = "Alibek"
	order1.BonusPoints = 10
	// order.IsReady = true // Undefined

	order2 := CoffeeOrder{
		CoffeeType:   "Cappuccino",
		Size:         "Medium",
		CustomerName: "Madi",
		BonusPoints:  5,
	}

	order3 := CoffeeOrder{
		CoffeeType:   "Americano",
		Size:         "Small",
		CustomerName: "Galym",
	}

	order4 := CoffeeOrder{"Espresso", "Small", "Rus", 15}

	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println(order3)
	fmt.Println(order4)
}
