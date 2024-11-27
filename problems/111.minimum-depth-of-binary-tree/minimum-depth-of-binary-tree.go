package main

import (
	"github.com/stolaar/leetcode/problems/utils"
)

type TreeNode = utils.TreeNode

func dfs(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	if node.Right == nil && node.Left == nil {
		return depth
	}

	if node.Right != nil && node.Left == nil {
		return dfs(node.Right, depth+1)
	}

	if node.Left != nil && node.Right == nil {
		return dfs(node.Left, depth+1)
	}

	return min(dfs(node.Right, depth+1), dfs(node.Left, depth+1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 1)
}

func main() {
	println(minDepth(utils.BuildTree("3,9,20,null,null,15,7")))         // 2
	println(minDepth(utils.BuildTree("2,null,3,null,4,null,5,null,6"))) // 5
}
