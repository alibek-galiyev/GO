package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType   string
	Size         string
	CustomerName string
	BonusPoints  int
}

type CoffeeMachine struct {
	Model          string
	Status         string
	OperationHours int
}

type CoffeeShop struct {
	Name string
}

type CoffeeType string

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

	fmt.Println(order1)
	printReceiptCoffee(order1)
	fmt.Println(order1)

	espressoMachine := CoffeeMachine{
		Model:          "Extra class espresso machine",
		Status:         "Operational",
		OperationHours: 75,
	}

	fmt.Println(espressoMachine)
	markIsOutOfService(&espressoMachine)
	fmt.Println(espressoMachine)

	myShop := CoffeeShop{
		Name: "Brew&Beans",
	}
	myShop.Greet()

	var myCoffee CoffeeType = "Espresso"
	myCoffee.Describe()

	orderStruct := CoffeeOrder{
		CoffeeType:   "Espresso",
		CustomerName: "Alibek",
		BonusPoints:  20,
		Size:         "Large",
	}

	orderMap := map[string]string{
		"CoffeeType":   "Espresso",
		"CustomerName": "Alibek",
		"Size":         "Large",
	}

	fmt.Println("-----------USING STRUCT-----------")
	fmt.Println("Customer", orderStruct.CustomerName)
	fmt.Println("Coffee Size", orderStruct.Size)
	fmt.Println()
	fmt.Println("-----------USING MAP-----------")
	fmt.Println("Customer", orderMap["CustomerName"])
	fmt.Println("Customer", orderMap["Size"])

}

func (s CoffeeType) Describe() {
	fmt.Println("This is the delicious", s)
}

func (shop *CoffeeShop) Greet() {
	fmt.Println("Welcome to", shop.Name)
}

func printReceiptCoffee(order CoffeeOrder) {
	order.Size = "Small"
	fmt.Println("Order size in the function:", order.Size)
}

func markIsOutOfService(machine *CoffeeMachine) {
	machine.Status = "Out of service"
	fmt.Println("Machine status changed to:", machine.Status)
}
