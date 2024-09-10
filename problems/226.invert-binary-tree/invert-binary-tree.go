package main

import (
	"github.com/stolaar/leetcode/problems/utils"
)

type TreeNode = utils.TreeNode

func dfs(current *TreeNode, left *TreeNode, right *TreeNode) {
	if left == nil && right == nil {
		return
	}

	if left == nil {
		current.Left = current.Right
		current.Right = nil

		dfs(right, right.Left, right.Right)
		return
	}

	if right == nil {
		current.Left, current.Right = current.Right, current.Left
		dfs(left, left.Left, left.Right)
		return
	}

	tmpl, tmpr := *left, *right
	*left, *right = tmpr, tmpl

	dfs(left, left.Left, left.Right)
	dfs(right, right.Left, right.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	current := root
	dfs(current, root.Left, root.Right)

	return root
}

func main() {
	utils.PrintTree(invertTree(utils.BuildTree("4,2,7,1,3,6,9")), 0)
	println()
	println()
	println()
	utils.PrintTree(invertTree(utils.BuildTree("2,1,3")), 0)
	println()
	println()
	println()
	utils.PrintTree(invertTree(utils.BuildTree("1,2")), 0)
}
