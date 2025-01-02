package main

import (
	"fmt"
)

func isVowel(r byte) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func vowelStrings(words []string, queries [][]int) []int {
	prefixSum := make([]int, len(words))

	count := 0

	for idx, word := range words {
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			count++
		}
		prefixSum[idx] = count
	}

	result := make([]int, len(queries))

	for idx, query := range queries {
		li, ri := query[0], query[1]

		if li-1 < 0 {
			result[idx] = prefixSum[ri]
			continue
		}

		result[idx] = prefixSum[ri] - prefixSum[li-1]
	}

	return result
}

func main() {
	fmt.Println(vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}))
}
