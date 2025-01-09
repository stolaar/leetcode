package main

import "strings"

func prefixCount(words []string, pref string) int {
	count := 0

	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}

	return count
}

func main() {
	println(prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
	println(prefixCount([]string{"leetcode", "win", "loops", "success"}, "code"))
}
