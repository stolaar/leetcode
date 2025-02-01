package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
func removeNeighbours(grid [][]byte, i int, j int) [][]byte {
	if j-1 >= 0 {
		if string(grid[i][j-1]) == "1" {
			grid[i][j-1] = []byte("0")[0]
			grid = removeNeighbours(grid, i, j-1)
		}
	}

	if j+1 < len(grid[i]) {
		if string(grid[i][j+1]) == "1" {
			grid[i][j+1] = []byte("0")[0]
			grid = removeNeighbours(grid, i, j+1)
		}
	}

	if i-1 >= 0 {
		if string(grid[i-1][j]) == "1" {
			grid[i-1][j] = []byte("0")[0]
			grid = removeNeighbours(grid, i-1, j)
		}
	}

	if i+1 < len(grid) {
		if string(grid[i+1][j]) == "1" {
			grid[i+1][j] = []byte("0")[0]
			grid = removeNeighbours(grid, i+1, j)
		}
	}

	return grid
}

// @lc code=start
func numIslands(grid [][]byte) int {
	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if string(grid[i][j]) == "1" {
				sum += 1
				grid = removeNeighbours(grid, i, j)
			}
		}
	}

	return sum
}


func main() {
	test := [][]byte{
		[]byte("111"),
		[]byte("010"),
		[]byte("111"),
	}

	test2 := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}

	test3 := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}

	test4 := [][]byte{
		[]byte("10111"),
		[]byte("10101"),
		[]byte("11101"),
	}

	result := numIslands(test)
	result2 := numIslands(test2)
	result3 := numIslands(test3)
	result4 := numIslands(test4)

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}

// @lc code=end
