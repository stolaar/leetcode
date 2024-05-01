package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 */
func sumOfLeftLeaves(root *TreeNode) int {
	var sum int = 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum = root.Left.Val
		} else {
			sum += sumOfLeftLeaves(root.Left)
		}
	}

	if root.Right != nil {
		return sum + sumOfLeftLeaves(root.Right)
	}

	return sum
}

// @lc code=end

func main() {
	fmt.Println(sumOfLeftLeaves(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}))
}
