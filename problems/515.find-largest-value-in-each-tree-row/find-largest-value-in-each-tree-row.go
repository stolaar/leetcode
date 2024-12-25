package main

import (
	"fmt"
	"slices"

	"github.com/stolaar/leetcode/problems/utils"
)

type TreeNode = utils.TreeNode

func bfs(left, right *TreeNode, level int, levels map[int][]int) {
	if _, ok := levels[level]; !ok && (left != nil || right != nil) {
		levels[level] = []int{}
	}

	if left != nil {
		levels[level] = append(levels[level], left.Val)
		bfs(left.Left, left.Right, level+1, levels)
	}

	if right != nil {
		levels[level] = append(levels[level], right.Val)
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
	levels := map[int][]int{}

	if root == nil {
		return []int{}
	}

	bfs(root.Left, root.Right, 1, levels)

	result := make([]int, len(levels)+1)

	result[0] = root.Val

	for level, values := range levels {
		result[level] = slices.Max(values)
	}

	return result
}

func main() {
	fmt.Println(largestValues(utils.BuildTree("1,3,2,5,3,null,9")))
	fmt.Println(largestValues(utils.BuildTree("1,2,3")))
}
