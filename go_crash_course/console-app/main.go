package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Print("-> ")

	// 	userInput, _ := reader.ReadString('\n')

	// 	userInput = strings.ReplaceAll(userInput, "\n", "")

	// 	if userInput == "quit" {
	// 		break
	// 	} else {
	// 		fmt.Println(userInput)
	// 	}
	// }

	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffes := make(map[int]string)
	coffes[1] = "Cappucino"
	coffes[2] = "Latte"
	coffes[3] = "Americano"
	coffes[4] = "Mocha"
	coffes[5] = "Macchiato"
	coffes[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))

		fmt.Println(fmt.Sprintf("You chose %s", coffes[i]))
	}

	fmt.Println("Program exiting ...")
}
