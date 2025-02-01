package main

import (
	"fmt"
	"strings"
	"time"
)

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
func isPalindrome(s string) bool {
	if len(s) == 2 {
		return s[0] == s[1]
	}
	var r string
	half := len(s) / 2
	first := s[0:half]

	for i := len(s) - 1; i >= half+1; i-- {
		r = r + string(s[i])
	}

	return first == r
}

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	start := 0
	end := 1

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return string(s[0])
	}

	maxLength := 0
	maxLengthRange := []int{0, 0}
	var result string
	streak := 0

	for end < len(s) && start < end {
		streak = end - start

		if len(result) > 0 && strings.HasPrefix(s[start:], result) {
			start = start + len(result)
			result = result + string(s[start+1]) + result
			maxLength += maxLength + 1
			end = start + 1
			continue
		}
		if isPalindrome(s[start : end+1]) {
			if streak > maxLength {
				maxLength = streak
				result = s[start : end+1]
				start = end + 1
				end = end + streak + 1
			}
			continue
		}

		found := false

		if streak > maxLength {
			i := start
			for i < end {
				if isPalindrome(s[i : end+1]) {
					if end-i > maxLength {
						maxLength = end - i
						result = s[i : end+1]
						start = end + 1
						end = end + streak
						break
					}
					end++
					continue
				}
				i++
			}
		}

		if found {
			continue
		}

		end++
	}

	if streak > maxLength {
		maxLength = streak
	}

	return s[maxLengthRange[0] : maxLengthRange[1]+1]

}

// @lc code=end

func main() {
	// fmt.Println("Longest", longestPalindrome("cbbd"))
	// fmt.Println("Longest", longestPalindrome("ac"))
	// fmt.Println("Longest", longestPalindrome("aa"))
	// fmt.Println("Longest", longestPalindrome("abc"))
	// fmt.Println("Longest", longestPalindrome("ccc"))
	// fmt.Println("Longest", longestPalindrome("cccc"))
	// fmt.Println("Longest", longestPalindrome("ana"))
	fmt.Println("Longest", longestPalindrome("bbabanadeifieddebes"))
	start := time.Now()
	fmt.Println("Longest", longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
	elapsed := time.Since(start)
	fmt.Printf("Exec time %s", elapsed)
}
