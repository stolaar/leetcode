package main

func maxScore(s string) int {
	ones, zeros := 0, 0

	result := 0

	for _, val := range s {
		if val == '0' {
			zeros++
			continue
		}
	}

	if zeros == 0 || zeros == len(s) {
		return len(s) - 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' && i > 0 {
			ones++

			if ones+zeros > result {
				result = ones + zeros
			}
			continue
		}

		if i == len(s)-1 {
			zeros--
			continue
		}

		if ones+zeros > result {
			result = ones + zeros
		}
		zeros--
	}

	return result
}

func main() {
	println(maxScore("011101"))
	println(maxScore("00111"))
	println(maxScore("1111"))
	println(maxScore("00"))
	println(maxScore("010"))
	println(maxScore("010"))
}
