package main

import "fmt"

func wordSubsets(words1 []string, words2 []string) []string {
	result := []string{}

	w1map, w2map := make([]map[rune]int, len(words1)), make([]map[rune]int, len(words2))
	maxw2len := 0

	w2dict := map[rune]int{}

	for i := 0; i < len(words1) || i < len(words2); i++ {
		if i < len(words1) {
			for _, r := range words1[i] {
				if w1map[i] == nil {
					w1map[i] = map[rune]int{r: 1}
					continue
				}
				w1map[i][r]++
			}
		}

		if i < len(words2) {
			if len(words2[i]) > maxw2len {
				maxw2len = len(words2[i])
			}

			for _, r := range words2[i] {
				if w2map[i] == nil {
					w2map[i] = map[rune]int{r: 1}
					continue
				}
				w2map[i][r]++
			}

			for k, v := range w2map[i] {
				w2dict[k] = max(w2dict[k], v)
			}
		}
	}

	for idx, word := range words1 {
		valid := true

		if len(word) < maxw2len {
			continue
		}

		for k, v := range w2dict {
			if w1map[idx][k] < v {
				valid = false
				break
			}
		}

		if valid {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))
}
