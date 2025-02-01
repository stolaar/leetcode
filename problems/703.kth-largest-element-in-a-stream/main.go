package main

import (
	"slices"
)

type KthLargest struct {
	K    int
	Nums []int
}

func Constructor(k int, nums []int) KthLargest {
	slices.SortFunc(nums, func(a int, b int) int {
		return b - a
	})

	return KthLargest{
		Nums: nums,
		K:    k,
	}
}

func (kth *KthLargest) Add(val int) int {
	i := len(kth.Nums)
	for idx, num := range kth.Nums {
		if val >= num {
			i = idx
			break
		}
	}

	if i == len(kth.Nums) {
		kth.Nums = append(kth.Nums, val)
	} else {
		kth.Nums = append(kth.Nums[:i+1], kth.Nums[i:]...)
		kth.Nums[i] = val
	}

	return kth.Nums[kth.K-1]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	println(obj.Add(3))
	println(obj.Add(5))
	println(obj.Add(10))
	println(obj.Add(9))
	println(obj.Add(4))
}
