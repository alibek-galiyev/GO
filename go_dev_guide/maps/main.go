package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#2EF527",
		"white": "#ffffff",
	}

	// colors["red"] = "#ff0000"
	// colors["green"] = "#2EF527"

	// delete(colors, "red")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
