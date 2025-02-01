package main

var openCloseMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

var closeOpenMap = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	valid, opened := true, []rune{}

	for _, r := range s {
		if _, ok := openCloseMap[r]; ok {
			opened = append(opened, r)
			continue
		}

		if validOpening, ok := closeOpenMap[r]; ok {
			if len(opened) == 0 {
				valid = false
				break
			}
			lastOpened := opened[len(opened)-1]

			if lastOpened != validOpening {
				valid = false
				break
			}

			opened = opened[:len(opened)-1]
		}
	}

	if len(opened) > 0 {
		return false
	}

	return valid
}

func main() {
	println(isValid("(]"))
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(({[]})())[)"))
}
