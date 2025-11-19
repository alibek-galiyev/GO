package main

import (
	"myApp/packageone"
)

// var one = "One"
var myVar = "MyVar"

func main() {

	var blockVar = "BlockVar"

	packageone.PrintMe(myVar, blockVar)

	// fmt.Println(one)

	// myFunc()

	// newString := packageone.PublicVar
	// fmt.Println(newString)

	// packageone.Exported()
}

// func myFunc() {
// 	// var one = "the number one"

// 	fmt.Println(one)
// }
