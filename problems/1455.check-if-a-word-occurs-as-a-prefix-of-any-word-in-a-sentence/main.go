package main

import (
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Fields(sentence)

	for idx, word := range words {
		if len(word) < len(searchWord) {
			continue
		}

		if strings.HasPrefix(word, searchWord) {
			return idx + 1
		}
	}

	return -1
}

func main() {
	println(isPrefixOfWord("i love eating burger", "burg"))
	println(isPrefixOfWord("i am tired", "you"))
}
