package main

import (
	"math"
)

func jump(nums []int) int {
	if len(nums) <= 2 {
		return len(nums) - 1
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

	return dist[0]
}

func main() {
	println(jump([]int{2, 3, 1, 1, 4}))
	println(jump([]int{2, 3, 0, 1, 4}))
	println(jump([]int{5, 3, 1, 8, 4, 2, 1}))
	println(jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}
