package main

import "sort"

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	result := 0

	for idx, val := range heights {
		if val != expected[idx] {
			result++
		}
	}

	return result
}

func main() {
	println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
