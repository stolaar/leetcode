package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	m1, m2 := map[string]int{}, map[string]int{}

	w1, w2 := strings.Split(s1, " "), strings.Split(s2, " ")
	i, i1 := 0, 0

	for i < len(w1) || i1 < len(w2) {
		if i < len(w1) {
			m1[w1[i]] += 1
		}

		if i1 < len(w2) {
			m2[w2[i1]] += 1
		}
		i++
		i1++
	}

	result := []string{}

	for key, val := range m1 {
		if m2[key] == 0 && val == 1 {
			result = append(result, key)
		}
	}
	for key, val := range m2 {
		if m1[key] == 0 && val == 1 {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	fmt.Println(uncommonFromSentences("this is sweet", "this is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}
