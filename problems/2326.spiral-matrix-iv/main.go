package main

import (
	"fmt"

	"github.com/stolaar/leetcode/problems/utils"
)

type ListNode = utils.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}

	directions := map[string][]int{
		"right": {0, 1, n - 1},
		"down":  {1, 0, m - 1},
		"left":  {0, -1, 0},
		"up":    {-1, 0, 1},
	}

	position := []int{0, 0}
	currentDirection := "right"

	current := head
	for current != nil {
		i, j := position[0], position[1]

		matrix[i][j] = current.Val

		current = current.Next
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

	return matrix
}

func main() {
	fmt.Println(spiralMatrix(3, 5, utils.CreateLinkedList([]int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0})))
	fmt.Println(spiralMatrix(1, 4, utils.CreateLinkedList([]int{0, 1, 2})))
	// fmt.Println(spiralMatrix(8, 7, utils.CreateLinkedList([]int{405, 341, 910, 645, 954, 915, 447, 924, 263, 350, 472, 1, 481, 972, 807, 970, 898, 525, 318, 620, 21, 922, 231, 192, 659, 976, 153, 130, 273, 997, 889, 488, 528, 832, 768, 444, 894, 682, 116, 569, 305, 112, 259, 810, 898, 831, 675, 165, 224, 367, 959, 783, 477, 974})))
}

