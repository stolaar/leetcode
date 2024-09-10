package main

import "github.com/stolaar/leetcode/problems/utils"

type TreeNode = utils.TreeNode

func dfs(root *TreeNode, result int) int {
	if root == nil {
		return result
	}
	result += 1

	l := dfs(root.Left, result)
	r := dfs(root.Right, result)

	return max(l, r)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func main() {
	println(maxDepth(utils.BuildTree("3,9,20,null,null,15,7")))
	println(maxDepth(utils.BuildTree("1,null,2")))
	println(maxDepth(utils.BuildTree("1,1,2,null,null,2,2,3,3,4,4,5,5")))
}
