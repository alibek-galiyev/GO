package main

import (
	"fmt"

	validparentheses "main.go/020_valid_parentheses"
)

func main() {
	s := "({[({[({[())]})]})]}"
	fmt.Println(validparentheses.IsValid(s))
}
