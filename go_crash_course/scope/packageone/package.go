package packageone

import "fmt"

var privateVar = "Private"
var PublicVar = "Public"
var PackageVar = "PackageVar"

func notExported() {

}

func Exported() {

}

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar)
	fmt.Println()
	fmt.Println(blockVar)
	fmt.Println()
	fmt.Println(PackageVar)
	fmt.Println()
	fmt.Println(privateVar)
}
