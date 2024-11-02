package main

import "strings"

func isCircularSentence(sentence string) bool {
	words, valid := strings.Split(sentence, " "), true

	if len(words) == 1 {
		return words[0][0] == words[0][len(words[0])-1]
	}

	for i := 1; i < len(words); i++ {
		pword, word := words[i-1], words[i]

		if pword[len(pword)-1] != word[0] {
			valid = false
			break
		}
	}

	if !valid {
		return false
	}

	lwidx := len(words) - 1
	lcidx := len(words[lwidx]) - 1
	return words[0][0] == words[lwidx][lcidx]
}

func main() {
	println(isCircularSentence("Hello oligarh"))
	println(isCircularSentence("leetcode exercises sound delightful"))
	println(isCircularSentence("eetcode"))
	println(isCircularSentence("Leetcode is cool"))
}
