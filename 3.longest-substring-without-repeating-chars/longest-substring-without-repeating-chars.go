package main

import "fmt"

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	start := 0
	end := 0

	maxLength := 0
	streak := 0

	for end < len(s) {
		i := start
		current := s[end]

		streak = end - start + 1

		continueChecking := true

		for i < end {
			if s[i] == current {
				if streak > maxLength {
					maxLength = streak - 1
				}
				start = i + 1
				continueChecking = false
				break
			}
			i++
		}

		if !continueChecking {
			continue
		}

		end++
	}

	if streak > maxLength {
		maxLength = streak
	}

	return maxLength
}

// @lc code=end
func main() {
	result := lengthOfLongestSubstring("abcabbacde")

	fmt.Print(result)
}
