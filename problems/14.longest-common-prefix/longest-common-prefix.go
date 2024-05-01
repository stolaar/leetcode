package main

import "fmt"

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start

func updatePrefixCount(prefixes map[string]int, str string) {
	if prefixes[str] < 1 {
		prefixes[str] = 1
	} else {
		prefixes[str] += 1
	}

}

func longestCommonPrefix(strs []string) string {
	prefixes := make(map[string]int)
	var longestPrefix string = ""

	for _, str := range strs {
		if len(str) == 0 {
			continue
		}
		if prefixes[string(str[0])] < 1 {
			prefixes[string(str[0])] = 1
		} else {
			prefixes[string(str[0])] += 1
		}
		start := 0
		end := 2

		for end <= len(str) {
			prefix := string(str[start:end])
			if prefixes[prefix] < 1 {
				prefixes[prefix] = 1
			} else {
				prefixes[prefix] += 1
			}
			end++
		}
	}

	for prefix, prefixLength := range prefixes {
		if prefixLength == len(strs) {
			if len(prefix) > len(longestPrefix) {
				longestPrefix = prefix
			}
		}
	}

	return longestPrefix
}

// @lc code=end

func main() {
	result := longestCommonPrefix([]string{"flower", "flow", "flight"})

	fmt.Println(result)
}
