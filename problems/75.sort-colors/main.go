package main

import "fmt"

func sortColors(nums []int) {
	s, n, p := 0, len(nums)-1, 0

	for p < 3 {
		for s < n {
			if nums[s] == p {
				s++
				continue
			}
			if nums[n] == p {
				nums[s], nums[n] = nums[n], nums[s]
				s++
				continue
			}
			n--
		}
		n = len(nums) - 1
		p++
	}
}

func main() {
	colors1 := []int{2, 0, 2, 1, 1, 0}
	colors2 := []int{0, 1, 2}
	sortColors(colors1)
	sortColors(colors2)
	fmt.Println(colors1)
	fmt.Println(colors2)
}
