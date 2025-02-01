package main

func countConsistentStrings(allowed string, words []string) int {
	m := make(map[rune]bool)

	for _, r := range allowed {
		m[r] = true
	}

	count := 0
	for _, word := range words {
		ok := true
		for _, wr := range word {
			if !m[wr] {
				ok = false
				break
			}
		}
		if ok {
			count += 1
		}
	}

	return count
}

func main() {
	println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
}
