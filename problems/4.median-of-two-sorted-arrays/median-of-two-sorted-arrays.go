package main

type medianPos struct {
	single bool
	x      int
	x2     int
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		x := len(nums2) / 2
		if len(nums2)%2 == 0 {
			n1, n2 := x-1, x

			return float64((nums2[n1] + nums2[n2])) / 2
		}

		return float64(nums2[x])
	}
	if len(nums2) == 0 {
		x := len(nums1) / 2
		if len(nums1)%2 == 0 {
			n1, n2 := x-1, x

			return float64((nums1[n1] + nums1[n2])) / 2
		}

		return float64(nums1[x])
	}

	ni1, ni2, len1, len2 := 0, 0, len(nums1), len(nums2)
	totalLen := len1 + len2
	var pos medianPos

	m := totalLen / 2
	if totalLen%2 == 0 {
		pos = medianPos{false, m - 1, m}
	} else {
		pos = medianPos{true, m, -1}
	}
	merged := []int{}
	var result float64

	for ni1 < len1 || ni2 < len2 {
		if pos.single && len(merged)-1 >= pos.x {
			result = float64(merged[pos.x])
			break
		}

		if !pos.single && len(merged)-1 >= pos.x2 {
			result = float64((merged[pos.x] + merged[pos.x2])) / 2
			break
		}
		if ni1 == len1 {
			merged = append(merged, nums2[ni2])
			ni2++
			continue
		}
		if ni2 == len2 {
			merged = append(merged, nums1[ni1])
			ni1++
			continue
		}

		if nums1[ni1] < nums2[ni2] {
			merged = append(merged, nums1[ni1])
			ni1++
			continue
		}
		if nums2[ni2] < nums1[ni1] {
			merged = append(merged, nums2[ni2])
			ni2++
			continue
		}

		if nums1[ni1] == nums2[ni2] {
			merged = append(merged, nums1[ni1], nums2[ni2])
			ni1++
			ni2++
			continue
		}
	}
	if pos.single && len(merged)-1 >= pos.x {
		result = float64(merged[pos.x])
	}

	if !pos.single && len(merged)-1 >= pos.x2 {
		result = float64((merged[pos.x] + merged[pos.x2])) / 2
	}

	return result
}

func main() {
	println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	println(findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 4, 4})) // 3.0000
}
