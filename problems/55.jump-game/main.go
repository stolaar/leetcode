package main

import "math"

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	dist := make([]int, len(nums))

	dist[len(nums)-1] = 0

	for i := len(nums) - 2; i >= 0; i-- {
		minDist := math.MaxInt64

		if nums[i] == 0 {
			dist[i] = -1
			continue
		}

		for j := i + 1; j <= min(len(nums)-1, i+nums[i]); j++ {
			if dist[j]+1 < minDist && dist[j] >= 0 {
				minDist = dist[j] + 1
			}
		}

		if minDist == math.MaxInt64 {
			dist[i] = -1
		} else {
			dist[i] = minDist
		}

	}

	return dist[0] > 0
}

func main() {
	println(canJump([]int{2, 3, 1, 1, 4}))
	println(canJump([]int{3, 2, 1, 0, 4}))
}
