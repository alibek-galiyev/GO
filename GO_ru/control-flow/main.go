package main

import "fmt"

func main() {
	orderTotal := 15.0
	// minRequiredOrderAmount := 10

	if orderTotal > 10 {
		fmt.Println("You got a free cookie for that order!")
	}

	orderTotal = 7.50
	if orderTotal > 10 {
		fmt.Println("You got a free cookie for that order!")
	} else if orderTotal > 5 {
		fmt.Println("You got a 50% discount!")
	} else {
		fmt.Println("You didn't get a free cookie or discount for that order!")
	}

	hour := 1630
	isMember := true
	orderAmount := 13.50

	if hour >= 1500 {
		if hour <= 1700 {
			if isMember {
				if orderAmount > 10 {
					fmt.Println("You got 30% discount!")
				} else {
					fmt.Println("Not enough amount!")
				}
			} else {
				fmt.Println("Not a member!")
			}
		} else {
			fmt.Println("Not a happy hour!")
		}
	}

	if fullAmt := getAmtWithTax(100, .1); fullAmt >= 110 {
		fmt.Println("You can join coffee club")
	}

	day := "Saturday"

	switch day {
	case "Monday", "Tuesday":
		fmt.Println("Weekday Specia")
	case "Saturday", "Sunday":
		fmt.Println("Weekend Special")
	default:
		fmt.Println("Unknown day")
	}

}

func getAmtWithTax(amount, tax float64) float64 {
	return amount + amount*tax
}
