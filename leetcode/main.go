package main

import (
	"fmt"

	longestcommonprefix "main.go/014_longest_common_prefix"
)

func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestcommonprefix.LongestCommonPrefix(strs))
}
