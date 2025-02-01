package main

import "github.com/stolaar/leetcode/problems/utils"

type TreeNode = utils.TreeNode

func dfs(root *TreeNode, sum int, targetSum int) bool {
	if root == nil {
		return sum == targetSum
	}

	if root.Left == nil && root.Right == nil {
		return sum == targetSum
	}

	left, right := false, false

	if root.Left != nil {
		left = dfs(root.Left, root.Left.Val+sum, targetSum)
	}
	if root.Right != nil {
		right = dfs(root.Right, root.Right.Val+sum, targetSum)
	}
	return left || right
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, root.Val, targetSum)
}

func main() {
	println(hasPathSum(utils.BuildTree("5,4,8,11,null,13,4,7,2,null,null,null,1"), 22))
	println(hasPathSum(utils.BuildTree("1,2,3"), 5))
	println(hasPathSum(utils.BuildTree("1,2,3"), 3))
	println(hasPathSum(utils.BuildTree("1,2,3"), 4))
	println(hasPathSum(utils.BuildTree(""), 0))
	println(hasPathSum(utils.BuildTree("1,2"), 1))
}
