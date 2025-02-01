package main

import "fmt"

func canMakeSubsequence(str1 string, str2 string) bool {
	i, j := 0, 0

	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] || str1[i]+1 == str2[j] || str1[i] == 'z' && str2[j] == 'a' {
			j += 1
			i += 1
			continue
		}

		i += 1
	}

	return j == len(str2)
}

func main() {
	fmt.Println(canMakeSubsequence("zc", "ad"))
}
