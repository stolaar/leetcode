package main

import "fmt"

func findMissingAndRepeatedValues(grid [][]int) []int {
	n, dict := len(grid), make(map[int]int)

	for k := 1; k <= n*n; k++ {
		dict[k]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dict[grid[i][j]]--
		}
	}

	result := make([]int, 2)
	for key, val := range dict {
		if val == -1 {
			result[0] = key
		}

		if val == 1 {
			result[1] = key
		}
	}

	return result
}

func main() {
	fmt.Println(findMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}}))
}
