package main

import (
	"fmt"
	"slices"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	ranks, sorted := map[int]int{}, make([]int, len(arr))
	copy(sorted, arr)

	slices.Sort(sorted)

	rank := 1
	ranks[sorted[0]] = rank

	for i := 1; i < len(sorted); i++ {
		if sorted[i] > sorted[i-1] {
			rank += 1
		}
		ranks[sorted[i]] = rank
		continue
	}

	for idx, num := range arr {
		sorted[idx] = ranks[num]
	}

	return sorted
}

func main() {
	fmt.Println(arrayRankTransform([]int{4, 1, 2, 3, 7}))
}
