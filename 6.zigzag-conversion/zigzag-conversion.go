package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

func generateColumnString(s string, padTop int, padBottom int) string {
	return strings.Repeat(" ", padTop) + s + strings.Repeat(" ", padBottom)
}

func convertToZigzag(s string, numRows int) []string {
	zigzag := []string{}

	firstColumn := s[0:numRows]
	for _, char := range firstColumn {
		zigzag = append(zigzag, string(char))
	}

	oneCharColumns := numRows - 2
	oneCharCount := 0

	for i := numRows; i < len(s); i++ {
		if oneCharCount < oneCharColumns {
			padTop := oneCharColumns - oneCharCount
			padBottom := numRows - padTop - 1

			columnString := generateColumnString(string(s[i]), padTop, padBottom)

			for idx, char := range columnString {
				if string(char) != " " {
					zigzag[idx] = zigzag[idx] + string(char)
				}
			}
			oneCharCount += 1
			continue
		}

		oneCharCount = 0

		if i+numRows > len(s) {
			for idx, char := range s[i:] {
				zigzag[idx] = zigzag[idx] + string(char)
			}
			break
		}

		for idx, char := range s[i : i+numRows] {
			zigzag[idx] = zigzag[idx] + string(char)
		}

		i += numRows - 1
	}

	return zigzag
}

// @lc code=start
func convert(s string, numRows int) string {
	if len(s) <= numRows {
		return s
	}
	zigzag := convertToZigzag(s, numRows)

	return strings.Join(zigzag, "")
}

// @lc code=end

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
