package main

func longestMonotonicSubarray(nums []int) int {
	result, maxInc, maxDec := 1, 1, 1

	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		if nums[i] < nums[j] {
			maxInc++
			maxDec = 1
		}

		if nums[i] > nums[j] {
			maxDec++
			maxInc = 1
		}

		if nums[i] == nums[j] {
			maxDec = 1
			maxInc = 1
			continue
		}

		if maxInc > result || maxDec > result {
			result = max(maxInc, maxDec)
			continue
		}
	}

	return result
}

func main() {
	println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2}))
}
