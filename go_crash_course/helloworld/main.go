package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.ReplaceAll(userInput, "\r\n", "")
		userInput = strings.ReplaceAll(userInput, "\n", "")

		if userInput == "quit" {
			break
		} else {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}

	}

}
