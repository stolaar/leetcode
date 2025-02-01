package main

import (
	"fmt"
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(a int, b int) bool {
		return len(dictionary[a]) < len(dictionary[b])
	})

	words := strings.Split(sentence, " ")
	startingChars := make(map[byte][]string)
	result := []string{}

	for _, word := range dictionary {
		if len(startingChars[word[0]]) > 0 {
			startingChars[word[0]] = append(startingChars[word[0]], word)
			continue
		}
		startingChars[word[0]] = []string{word}
	}

	for _, word := range words {
		if len(startingChars[word[0]]) > 0 {
			found := false
			for _, dict := range startingChars[word[0]] {
				if strings.HasPrefix(word, dict) {
					result = append(result, dict)
					found = true
					break
				}
			}

			if found {
				continue
			}
		}

		result = append(result, word)
	}

	return strings.Join(result, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was ratted by the battery"))
	fmt.Println(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}
