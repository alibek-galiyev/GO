package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnsADog = readYesNo("Do you have a dog (y/n)?")

	fmt.Printf(
		"Your name is %s. You are %d years old. Your favourite number is %.2f. Status of dog: %t.\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwnsADog,
	)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.ReplaceAll(userInput, "\r\n", "")
		userInput = strings.ReplaceAll(userInput, "\n", "")

		if userInput == "" {
			fmt.Println("Please enter a value!")
		} else {
			return userInput
		}

	}

}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.ReplaceAll(userInput, "\r\n", "")
		userInput = strings.ReplaceAll(userInput, "\n", "")

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number!")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.ReplaceAll(userInput, "\r\n", "")
		userInput = strings.ReplaceAll(userInput, "\n", "")

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a number!")
		} else {
			return num
		}
	}
}

func readYesNo(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal()
		}

		if strings.ToLower(string(char)) == "y" || strings.ToLower(string(char)) == "n" {
			if char == 'n' || char == 'N' {
				return false
			} else {
				return true
			}
		} else {
			fmt.Println("Please enter 'y' or 'n'!")
		}
	}
}
