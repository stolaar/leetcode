package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	directions := map[string][]int{
		"right": {0, 1, len(matrix[0]) - 1},
		"down":  {1, 0, len(matrix) - 1},
		"left":  {0, -1, 0},
		"up":    {-1, 0, 1},
	}

	position := []int{0, 0}
	currentDirection := "right"

	for i2 := 0; i2 < len(matrix); i2++ {
		for j2 := 0; j2 < len(matrix[0]); j2++ {
			i, j := position[0], position[1]

			result = append(result, matrix[i][j])

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
	}

	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
