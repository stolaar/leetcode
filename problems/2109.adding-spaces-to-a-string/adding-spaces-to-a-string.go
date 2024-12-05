package main

import (
	"strings"
)

func addSpaces(s string, spaces []int) string {
	j := 0
	var result strings.Builder

	for i := 0; i < len(s); i++ {
		if j >= len(spaces) {
			result.WriteString(s[i:])
			break
		}

		if spaces[j] == i {
			result.WriteString(" ")
			j += 1
		}

		result.WriteByte(s[i])
	}

	return result.String()
}

func main() {
	println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
	println(addSpaces("icodeinpython", []int{1, 5, 7, 9}))
	println(addSpaces("spacing", []int{0, 1, 2, 3, 4, 5, 6}))
}
