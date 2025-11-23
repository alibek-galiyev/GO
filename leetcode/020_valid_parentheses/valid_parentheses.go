package validparentheses

func IsValid(s string) bool {
	var stack []rune
	var brackets map[rune]rune = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}
		if open, isClosing := brackets[ch]; isClosing {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
