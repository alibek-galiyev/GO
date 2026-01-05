package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range numbers {
		fmt.Printf("Element %d is %s\n", value, checkOdd(value))
	}
}

func checkOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
