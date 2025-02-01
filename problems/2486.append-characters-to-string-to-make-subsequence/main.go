package main

func appendCharacters(s string, t string) int {
  sp, tp, matching := 0,0,0


  for sp < len(s) && tp < len(t) {
    if s[sp] == t[tp] {
      matching++
      sp++
      tp++
      continue
    }

    sp++
  }


  return len(t) - matching
}

func main() {
  println(appendCharacters("coaching", "coding"))
  println(appendCharacters("abcde", "a"))
  println(appendCharacters("z", "abcde"))
}
