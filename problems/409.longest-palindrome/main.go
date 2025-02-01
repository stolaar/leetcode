package main

func longestPalindrome(s string) int {
  pairs := make(map[rune]int)
  streak := 0

  for _, char := range s {
    pairs[char]++

    if pairs[char] == 2 {
      streak += 2
      pairs[char] = 0
    }
  }

  f := false
  for _, count := range pairs {
    if count == 1 {
      f = true
      break
    }
  }

  if f {
    return streak + 1
  }

  return streak
}

func main() {
  println(longestPalindrome("abccccdd"))
  println(longestPalindrome("a"))
  println(longestPalindrome("aaaa"))
}
