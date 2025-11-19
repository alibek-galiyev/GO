package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	greet("Brew&Beans")
	fmt.Println(info("Alibek", "Cappuccino", 4.40))
	fmt.Println(calculator(3, 4, "+"))
	fmt.Println(calculator(1, 0, "/"))
	fmt.Println(circleParameters(4.66, pi))

	piEst := func() float64 {
		res := 0.0
		determinant := 1.0
		for range 100 {
			res += 1 / determinant
			res -= 1 / (determinant + 2)
			determinant += 4
		}
		return 4 * res
	}

	fmt.Println(circleParameters(4.66, piEst()))

	c1 := counter()
	for range 5 {
		fmt.Println(c1(2))
	}

}

func counter() func(increment int) int {
	c := 0
	return func(increment int) int {
		c += increment
		return c
	}
}

func greet(name string) {
	fmt.Println("Welcome to the Coffee Shop", name)
}

func info(name, coffeeType string, price float32) string {
	return fmt.Sprintf(
		"Customer with name %s has a coffee %s for price $%.2f",
		name,
		coffeeType,
		price,
	)
}

func calculator(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0.0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("unknown operation")
	}
}

func circleParameters(r, pi float64) (circleArea, circleCircumference float64) {
	circleArea = r * r * pi
	circleCircumference = 2.0 * pi * r
	return // naked return
}
