package main

import (
	"fmt"
	"strings"
)

func compressedString(word string) string {
	var comp strings.Builder

	i := 0

	for i < len(word) {
		lenght, j := 1, i+1

		for j < len(word) {
			if word[j] != word[i] {
				break
			}
			lenght += 1

			if lenght == 9 {
				break
			}

			j += 1
		}

		comp.WriteString(fmt.Sprintf("%d", lenght))
		comp.WriteByte(word[i])

		i += lenght
	}

	return comp.String()
}

func main() {
	println(compressedString("aaaaaaaaaaaaaabb"))
	println(compressedString("abcde"))
}
