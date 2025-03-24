package main

import (
	"math"
)

func longestNiceSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	longest, streak, left, right, curr := 1, 1, 0, 1, 0

	for right < len(nums) {
		if streak > longest {
			longest = streak
		}

		if longest == 30 {
			return 30
		}
		if nums[left]&nums[right] == 0 && curr&nums[right] == 0 {
			streak++
			curr |= nums[left] | nums[right]
			right++

			continue
		}

		if nums[right]&(curr^nums[left]) == 0 {
			curr ^= nums[left]
			left += 1
			right = int(math.Max(float64(left+1), float64(right)))

			streak = int(math.Max(float64(right-left), 1))
			continue
		}

		curr ^= nums[left]
		left++
		for left < right && nums[right]&(curr) != 0 {
			left++
			curr ^= nums[left]
		}

		if nums[left-1]&nums[right] == 0 {
			left--
		}

		if left == right {
			right += 1
		}

		streak = int(math.Max(float64(right-left), 1))
		curr = 0
	}

	if streak > longest {
		return streak
	}

	return longest
}

func main() {
	println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))                                                                                                                                                                                       // 3
	println(longestNiceSubarray([]int{744437702, 379056602, 145555074, 392756761, 560864007, 934981918, 113312475, 1090, 16384, 33, 217313281, 117883195, 978927664}))                                                                         // 3
	println(longestNiceSubarray([]int{135745088, 609245787, 16, 2048, 2097152}))                                                                                                                                                               // 3
	println(longestNiceSubarray([]int{84139415, 693324769, 614626365, 497710833, 615598711, 264, 65552, 50331652, 1, 1048576, 16384, 544, 270532608, 151813349, 221976871, 678178917, 845710321, 751376227, 331656525, 739558112, 267703680})) // 8
	println(longestNiceSubarray([]int{21, 2, 8, 23, 64, 128, 256, 512}))                                                                                                                                                                       // 6
}
