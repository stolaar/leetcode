package main

import "strings"

func isPalindrome(s string) bool {
	str := ""

	for _, r := range s {
		if (r >= 97 && r <= 122) || (r >= 48 && r <= 57) {
			str += string(r)
		}

		if r >= 65 && r <= 90 {
			str += strings.ToLower(string(r))
		}
	}

	if len(str) <= 1 {
		return true
	}

	l, r, m := 0, len(str)-1, len(str)/2

	for l < m {
		if str[l] != str[r] {
			return false
		}
		l += 1
		r -= 1
	}

	return true
}

func main() {
	println(isPalindrome("A man, a plan, a canal: Panama"))
	println(isPalindrome("race a car"))
	println(isPalindrome(" "))
	println(isPalindrome("0P"))
}
