package main

import "github.com/stolaar/leetcode/problems/utils"

type TreeNode = utils.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return left && right
}

func main() {
	println(isSameTree(utils.BuildTree("1,2,3"), utils.BuildTree("1,2,3")))
	println(isSameTree(utils.BuildTree("1,2,3"), utils.BuildTree("1,null,3")))
	println(isSameTree(utils.BuildTree(""), utils.BuildTree("")))
	println(isSameTree(utils.BuildTree("1"), utils.BuildTree("")))
	println(isSameTree(utils.BuildTree("1,2,3,4,5"), utils.BuildTree("1,2,3,4,5")))
	println(isSameTree(utils.BuildTree("1,null,3,4,5"), utils.BuildTree("1,null,3,4,5")))
}
