package romantointeger

var romanNumbers map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	res := 0
	for i := 0; i < len(s)-1; i++ {
		curr := romanNumbers[s[i]]
		next := romanNumbers[s[i+1]]
		if curr < next {
			res -= curr
		} else {
			res += curr
		}
	}
	res += romanNumbers[s[len(s)-1]]
	return res
}
