package main

import (
	"fmt"
	"strings"
	"time"
)

func findLongestWord(words []string) string {
  longest := words[0]

  for i:=1;i<len(words);i++ {
    if len(words[i])>len(longest) {
      longest = words[i]
    }
  }

  return longest
}

func commonChars(words []string) []string {
  result := []string{}

  longest := findLongestWord(words)

  for _, char := range longest {
    exists := true
    for idx, word := range words {
      if !strings.Contains(word, string(char)) {
        exists = false
        break
      } else {
        words[idx] = strings.Replace(word, string(char), "", 1)
      }
    }

    if exists {
      result = append(result, string(char))
    }
  }

  return result
}

func main() {
  start := time.Now()
  fmt.Println(commonChars([]string{"bella", "label", "roller"}))
  fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
  fmt.Println("end", time.Since(start))
}
