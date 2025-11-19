package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// one way - declare, then assign
	var firstNumber int
	firstNumber = rand.Intn(8) + 2

	// another way, declare type and name and assign value
	var secondNumber = rand.Intn(8) + 2

	// one step variable
	subtraction := rand.Intn(8) + 2

	gamePrompts(firstNumber, secondNumber, subtraction)

}

func gamePrompts(firstNumber, secondNumber, subtraction int) {

	answer := firstNumber*secondNumber - subtraction

	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("_____________________")
	fmt.Println("")

	fmt.Println("Think a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them an answer
	fmt.Println("The answer is", answer)

}
