package main

import "fmt"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numIdx := map[int]int{}
	result := []int{}

	for idx, num := range nums {
		corresponding := target - num

		if numIdx[corresponding] >= 1 {
			result = append(result, idx, numIdx[corresponding]-1)
			break
		}
		numIdx[num] = idx + 1
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
