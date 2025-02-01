package main

func waysToSplitArray(nums []int) int {
	prefixSum := make([]int, len(nums))

	prefixSum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	result, n := 0, len(nums)

	for i := 0; i < len(prefixSum)-1; i++ {
		rightSum := (prefixSum[n-1] - prefixSum[i])
		if prefixSum[i] >= rightSum && i < n-1 {
			result++
		}
	}

	return result
}

func main() {
	println(waysToSplitArray([]int{10, 4, -8, 7}))
	println(waysToSplitArray([]int{2, 3, 1, 0}))
}
