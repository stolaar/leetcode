package main

import (
	"fmt"
	"maps"
)

func findWordCommonChars(word string, charsMap map[rune]int, idx int) {
  wordCharsMap := make(map[rune]int)

  for _, char := range word {
    wordCharsMap[char]++
  }

  if idx == 0 {
    maps.Copy(charsMap, wordCharsMap)
    return
  }

  for char := range charsMap {
    if wordCharsMap[char] <= 0 {
      charsMap[char] = 0
      continue
    }

    charsMap[char] = min(charsMap[char], wordCharsMap[char])
  }

}

func commonChars(words []string) []string {
  charsMap := make(map[rune]int)
  result := []string{}

  for idx, word := range words {
      findWordCommonChars(word, charsMap, idx)
  }

  for char, count := range charsMap {
    for i:=0;i<count;i++ {
      result = append(result, string(char))
    }
  }

  return result
}

func main() {
  fmt.Println(commonChars([]string{"bella", "label", "roller"}))
  fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}
