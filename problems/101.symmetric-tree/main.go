package main

import "github.com/stolaar/leetcode/problems/utils"

type TreeNode = utils.TreeNode

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left, root.Right)
}

func main() {
	println(isSymmetric(utils.BuildTree("1,2,2,3,4,4,3")))
	println(isSymmetric(utils.BuildTree("1,2,2,null,3,null,3")))
}
