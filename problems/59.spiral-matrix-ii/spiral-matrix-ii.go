package main

import "fmt"

func generateMatrix(n int) [][]int {
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	directions := map[string][]int{
		"right": {0, 1, n - 1},
		"down":  {1, 0, n - 1},
		"left":  {0, -1, 0},
		"up":    {-1, 0, 1},
	}

	position := []int{0, 0}
	currentDirection := "right"
	current := 1

	for current <= n*n {
		i, j := position[0], position[1]

		result[i][j] = current
		current += 1

		direction := directions[currentDirection]

		if j == direction[2] && currentDirection == "right" {
			direction[2] -= 1
			currentDirection = "down"
			position[0], position[1] = i+directions[currentDirection][0], j+directions[currentDirection][1]
			continue
		}

		if i == direction[2] && currentDirection == "down" {
			direction[2] -= 1
			currentDirection = "left"
			position[0], position[1] = i+directions[currentDirection][0], j+directions[currentDirection][1]
			continue
		}

		if j == direction[2] && currentDirection == "left" {
			direction[2] += 1
			currentDirection = "up"
			position[0], position[1] = i+directions[currentDirection][0], j+directions[currentDirection][1]
			continue
		}

		if i == direction[2] && currentDirection == "up" {
			direction[2] += 1
			currentDirection = "right"
			position[0], position[1] = i+directions[currentDirection][0], j+directions[currentDirection][1]
			continue
		}

		position[0], position[1] = i+directions[currentDirection][0], j+directions[currentDirection][1]

	}

	return result
}

func main() {
	fmt.Println(generateMatrix(3))
}
