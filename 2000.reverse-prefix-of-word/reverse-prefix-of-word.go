package main

import "strings"

func reversePrefix(word string, ch byte) string {
  if !strings.Contains(word, string(ch)) {
    return word
  }

  prefix := ""

  for idx, char := range word {
    if byte(char) == ch {
      prefix = string(char) + prefix
      prefix += word[idx+1:]
      break
    }
    prefix = string(char) + prefix
  }


  return prefix
}

func main() {
  println(reversePrefix("abcd", 'z'))
  println(reversePrefix("abcdefd", 'd'))
}
