package main

import (
	"sort"

	"github.com/stolaar/leetcode/problems/utils"
)

type TreeNode = utils.TreeNode

func bfs(root *TreeNode, level int, result []int) []int {
	if root == nil {
		return result
	}

	if root.Left == nil && root.Right == nil {
		return result
	}

	levelSum := 0

	if root.Left != nil {
		levelSum += root.Left.Val
	}

	if root.Right != nil {
		levelSum += root.Right.Val
	}

	if len(result) == level {
		result = append(result, levelSum)
	} else {
		result[level] += levelSum
	}

	nextLevel := level + 1

	result = bfs(root.Left, nextLevel, result)
	result = bfs(root.Right, nextLevel, result)

	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levels := []int{root.Val}

	levels = bfs(root, 1, levels)

	sort.Ints(levels)

	if len(levels) < k {
		return -1
	}

	for i := len(levels) - 1; i >= 0; i-- {
		if k == 1 {
			return int64(levels[i])
		}

		k--
	}

	return int64(levels[0])
}

func main() {
	println(kthLargestLevelSum(utils.BuildTree("5,8,9,2,1,3,7,4,6"), 2))
	println(kthLargestLevelSum(utils.BuildTree("5,8,9,2,1,3,7"), 4))
}
