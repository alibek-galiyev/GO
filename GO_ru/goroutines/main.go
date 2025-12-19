package main

import (
	"fmt"
	"sync"
	"time"
)

func prepareDrink(barista string, drink string, c chan string) {
	fmt.Printf("Barista %s starting to prepare drink...\n", barista)
	time.Sleep(2 * time.Second)
	msg := fmt.Sprintf("Barista %s done preparing %s\n", barista, drink)
	c <- msg

}

func divideByZero(a int) {
	fmt.Println(10 / a)
}

func makeDrink(barista string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Barista %s: Starting to make a coffee...\n", barista)
	time.Sleep(2 * time.Second)
	fmt.Printf("Barista %s: Done!\n", barista)
}

func main() {
	orderChannel := make(chan string)

	var wg sync.WaitGroup

	fmt.Println("Coffee Shop opens")

	baristas := []string{"Aliba", "Madi", "Galym"}
	drinks := []string{"Latte", "Cappuccino", "Espresso"}

	for _, name := range baristas {
		wg.Add(1)
		go makeDrink(name, &wg)
	}

	wg.Wait()

	for i, name := range baristas {
		go prepareDrink(name, drinks[i], orderChannel)
	}

	for range baristas {
		msg := <-orderChannel
		fmt.Println(msg)
	}

	fmt.Println("All drinks are ready")
	fmt.Println("Coffee Shop closes")
}
