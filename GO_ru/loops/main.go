package main

import "fmt"

func main() {
	coffeeCupsQty := 10

	for i := 1; i <= coffeeCupsQty; i++ {
		fmt.Println("Value of the i:", i)
	}

	for i := coffeeCupsQty; i > 0; i-- {
		fmt.Printf("Coffee number: #%d\n", i)
	}

	tokens := 3

	for tokens > 0 {
		fmt.Println("1. Making yet another cup of coffee")
		tokens--
	}

	tokens = 5

	for {
		if tokens < 0 {
			break
		} else {
			fmt.Println("2. Making yet another cup of coffee")
			tokens--
		}
	}

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("---------------------Welcome to the brew&beans-------------------")
	fmt.Println("-----------------------------------------------------------------")

	for {
		var order string
		fmt.Println("Enter your coffee name (or type 'exit' to quit):")
		fmt.Scanln(&order)
		if order == "exit" {
			fmt.Println("Thank you. Good Bye!")
			break
		}
		if order == "" {
			fmt.Println("Please enter valid type")
			continue
		}
		fmt.Println("You ordered:", order)
	}

	dailyIncome := []float64{120.50, 98.75, 101.23, 3843.54, 34.3434, 45.56}
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	fmt.Println("Weekly Income Report:")
	var totalIncome float64

	for i, v := range dailyIncome {
		fmt.Printf("At %s income was: %.2f\n", daysOfWeek[i], v)
		totalIncome += v
	}
	fmt.Println("Total income for the week is", totalIncome)

	drinks := []string{"Latte", "Cappuccino", "Matcha", "Expired Milk", "Espresso"}

	for _, v := range drinks {
		if v == "Expired Milk" {
			continue
		}
		if v == "Matcha" {
			fmt.Println("Sold Out")
			break
		}
		fmt.Println("Preparing drink:", v)
	}
}
