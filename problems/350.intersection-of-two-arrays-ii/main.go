package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	occ1, occ2, i := map[int]int{}, map[int]int{}, 0
	intersectArr := []int{}

	for i < len(nums1) || i < len(nums2) {
		if i < len(nums1) {
			if occ2[nums1[i]] > 0 {
				occ2[nums1[i]]--
				intersectArr = append(intersectArr, nums1[i])
			} else {
				occ1[nums1[i]]++
			}
		}
		if i < len(nums2) {
			if occ1[nums2[i]] > 0 {
				occ1[nums2[i]]--
				intersectArr = append(intersectArr, nums2[i])
			} else {
				occ2[nums2[i]]++
			}
		}
		i++
	}

	for key, value := range occ1 {
		if occ2[key] > 0 {
			v := min(occ2[key], value)
			for i := 0; i < v; i++ {
				intersectArr = append(intersectArr, key)
			}
		}
	}

	return intersectArr
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 3}, []int{2, 2}))
}
