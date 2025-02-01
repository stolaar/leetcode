package main

/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

func checkSquarePerimeter(grid [][]int, i int, j int) int {
	sum := 0

	if j == 0 || j == len(grid[i]) - 1 {
		sum += 1
	}

	if i == 0 || i == len(grid) {
		sum += 1
	}

	if j - 1 >= 0 {
		if grid[i][j-1] == 0 {
			sum += 1
		}
	}

	if j + 1 < len(grid[i]) {
		if grid[i][j+1] == 0 {
			sum += 1
		}
	}

	if i - 1 >= 0 {
		if grid[i-1][j] == 0 {
			sum += 1
		}
	}

	if i + 1 < len(grid[i]) {
		if grid[i+1][j] == 0 {
			sum += 1
		}
	}
	return sum
}

// @lc code=start
func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			square := grid[i][j]
			if square == 0 {
				continue
			}

			perimeter += checkSquarePerimeter(grid, i, j)
			j++
		}

		i++
	}

	return perimeter
}

// @lc code=end
