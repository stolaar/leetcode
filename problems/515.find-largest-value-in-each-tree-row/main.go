package main

import (
	"fmt"
	"math"

	"github.com/stolaar/leetcode/problems/utils"
)

type TreeNode = utils.TreeNode

func bfs(left, right *TreeNode, level int, levels map[int]int) {
	if _, ok := levels[level]; !ok && (left != nil || right != nil) {
		levels[level] = math.MinInt64
	}

	if left != nil {
		if levels[level] < left.Val {
			levels[level] = left.Val
		}

		bfs(left.Left, left.Right, level+1, levels)
	}

	if right != nil {
		if levels[level] < right.Val {
			levels[level] = right.Val
		}
		bfs(right.Left, right.Right, level+1, levels)
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	levels := map[int]int{}

	if root == nil {
		return []int{}
	}

	bfs(root.Left, root.Right, 1, levels)

	result := make([]int, len(levels)+1)

	result[0] = root.Val

	for level, value := range levels {
		result[level] = value
	}

	return result
}

func main() {
	fmt.Println(largestValues(utils.BuildTree("1,3,2,5,3,null,9")))
	fmt.Println(largestValues(utils.BuildTree("1,2,3")))
}
