package main

import "fmt"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

func groupString(str string, groupMaps map[int]map[rune]int, result [][]string) (map[int]map[rune]int, [][]string) {
	charsCount := make(map[rune]int)
	existingIndex := -1

	for _, char := range str {
		if charsCount[char] <= 0 {
			charsCount[char] = 1
			continue
		}

		charsCount[char] += 1
	}

	for idx, group := range groupMaps {
		if len(charsCount) != len(group) {
			continue
		}
		equals := true

		for char := range group {
			if group[char] != charsCount[char] {
				equals = false
				break
			}
		}

		if equals {
			existingIndex = idx
			break
		}
	}

	if existingIndex < 0 {
		groupMaps[len(result)] = charsCount
		result = append(result, []string{str})

		return groupMaps, result
	}

	result[existingIndex] = append(result[existingIndex], str)

	return groupMaps, result
}

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	groupMaps := make(map[int]map[rune]int)
	result := [][]string{}

	for _, str := range strs {
		groupMaps, result = groupString(str, groupMaps, result)
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))

}
