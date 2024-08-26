package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	result := []int{}

	postorder := func(n *TreeNode) {}
	postorder = func(n *TreeNode) {
		if n == nil {
			return
		}

		postorder(n.Left)
		postorder(n.Right)

		result = append(result, n.Val)
	}

	postorder(root)

	return result
}

func main() {
	fmt.Println(postorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))

	fmt.Println(postorderTraversal(&TreeNode{
		Val: 1,
	}))
}
