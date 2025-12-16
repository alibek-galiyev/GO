package main

import (
	"fmt"
	"time"
)

func divideByZero(a int) {
	fmt.Println(10 / a)
}

func makeDrink(barista string) {
	fmt.Printf("Barista %s: Starting to make a coffee...\n", barista)
	time.Sleep(2 * time.Second)
	fmt.Printf("Barista %s: Done!\n", barista)
}

func main() {
	fmt.Println("Coffee Shop opens")
	go makeDrink("Aliba")
	go makeDrink("Madi")
	go makeDrink("Alex")
	time.Sleep(3 * time.Second)
	fmt.Println("Coffee Shop closes")
}
