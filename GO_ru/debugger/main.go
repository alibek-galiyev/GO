package main

import (
	"errors"
	"fmt"
	"os"
)

type CoffeeError string

func (e CoffeeError) Error() string {
	return string(e)
}

func readFiles() {
	file, err := os.Open("coffee_orders.txt")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File not exist!")
		} else {
			fmt.Println("General file opening error!")
		}
		return
	}

	fmt.Println("Successfully accessed file ...")
	fmt.Println("Successfully accessed file: ", file.Name())
}

func closeShop() {
	fmt.Println("Closing the shop!")
}

func DispenseCoffee(coffeeAmount int, cups int) {
	// if cups <= 0 {
	// 	fmt.Println("Error: Zero Division error")
	// 	return
	// }

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}

	}()
	fmt.Printf("Dispensing %d gramms of coffee into %d cups ...\n", coffeeAmount, cups)
	amountPerCup := coffeeAmount / cups
	fmt.Printf("Each cup gets %d gramms of coffee\n", amountPerCup)
}

func main() {
	fmt.Println("Starting machine ...")

	DispenseCoffee(750, 200)

	fmt.Println("Coffee machine is still running ...")
	DispenseCoffee(340, 0)

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	defer closeShop()
	fmt.Println("Opening the coffee shop ...")
	fmt.Println("Serving a customer ...")

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	defer func() {
		fmt.Println("Cleaning a coffee machine ...")
		fmt.Println("Suspending a coffee machine ...")
	}()
	fmt.Println("Brewing fresh cup of coffee ...")

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	fmt.Println("Starting machine ...")

	DispenseCoffee(750, 200)

	fmt.Println("Coffee machine is still running ...")
	DispenseCoffee(340, 0)

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	readFiles()

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	var err error

	err = fmt.Errorf("Some error")

	if err == nil {
		fmt.Println("There's no error!")
	} else {
		fmt.Println("Error occured!", err)
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	// var error error
	// error = CoffeeError("No coffee beans loaded!")

	// if error != nil {
	// 	fmt.Println("Error:", error)
	// 	panic(error)
	// }

	// panic(error)

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	msg, errr := ServeDrink("Espresso")
	if errr != nil {
		fmt.Println("Serving failed", errr)
	} else {
		fmt.Println(msg)
	}

	msg, errr = ServeDrink("Latte")

	if errr != nil {
		fmt.Println("Serving failed", errr)
	} else {
		fmt.Println(msg)
	}

	msg, errr = ServeDrink("Tea")

	if errr != nil {
		fmt.Println("Serving failed", errr)
	} else {
		fmt.Println(msg)
	}

	fmt.Println("Stock after servings:", stock)

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")

	temps := []int{85, 95, 105}

	for _, t := range temps {
		err := CheckTemperature(t)
		if err != nil {
			fmt.Println("Temperature check failed:", err)
		} else {
			fmt.Println("Temperature is OK:", t)
		}
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")
}

var stock = map[string]int{
	"Espresso":   5,
	"Latte":      0,
	"Cappuccino": 3,
}

type OutOfStockError struct {
	Item string
}

func (e OutOfStockError) Error() string {
	return fmt.Sprintf("%s is out of stock\n", e.Item)
}

func ServeDrink(item string) (string, error) {
	qty := stock[item]
	if qty == 0 {
		return "", OutOfStockError{Item: item}
	}
	stock[item]--
	return fmt.Sprintf("Here's your freshly brewed %s", item), nil
}

func CheckTemperature(temp int) error {
	if temp > 100 {
		return errors.New("Temperature too high!")
	}
	if temp > 90 {
		return fmt.Errorf("Machine overheated at %d grad C", temp)
	}
	return nil
}
