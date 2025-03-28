package main

func minChanges(s string) int {
	count := 0

	for i := 0; i < len(s)-1; i += 2 {
		if s[i] != s[i+1] {
			count += 1
		}
	}

	return count
}

func main() {
	println(minChanges("1001"))
	println(minChanges("0000"))
	println(minChanges("10"))
}
