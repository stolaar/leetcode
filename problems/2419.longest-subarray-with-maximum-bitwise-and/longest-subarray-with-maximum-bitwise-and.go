package main

func longestSubarray(nums []int) int {
	result, currentResult, maxval, prev := 0, 0, nums[0], nums[0]

	for _, num := range nums {
		if num >= maxval && prev != num {

			if maxval == num {
				result = max(result, currentResult)
			}
			maxval = num

			if num&prev == num {
				currentResult = 2
			} else {
				currentResult = 1
			}

			prev = num
			continue
		}

		if num == maxval && prev == num {
			currentResult += 1
			continue
		}

		if num&prev == maxval {
			currentResult += 1
			prev = num
			continue
		}

		prev = num
	}

	return max(currentResult, result)
}

func main() {
	// println(longestSubarray([]int{1, 2, 3, 3, 2, 2}))
	// println(longestSubarray([]int{1, 2, 3, 4}))
	// println(longestSubarray([]int{311155, 311155, 311155, 311155, 311155, 311155, 311155, 311155, 201191, 311155}))
	println(longestSubarray([]int{395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 153490, 330001, 330001, 330001, 330001, 330001, 330001, 330001, 37284, 470030, 470030, 470030, 470030, 470030, 470030, 156542, 226743}))
}
