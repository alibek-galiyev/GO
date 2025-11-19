package palindromenumber

import "strconv"

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	n := len(s) / 2
	for i := 0; i < n; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
