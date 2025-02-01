package main

import "strings"

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}

	var result strings.Builder
	result.WriteByte(s[0])
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if count <= 2 {
				result.WriteByte(s[i])
			}
		} else {
			count = 1
			result.WriteByte(s[i])
		}
	}

	return result.String()
}

func main() {
	println(makeFancyString("leeetcode"))
	println(makeFancyString("uuups"))
	println(makeFancyString("nooopeee"))
}
