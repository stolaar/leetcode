package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func dfs(node *Node, values []int) []int {
	if node == nil {
		return values
	}
	for _, n := range node.Children {
		values = dfs(n, values)
	}

	values = append(values, node.Val)
	return values
}

func postorder(root *Node) []int {
	result := []int{}

	result = dfs(root, result)

	return result
}

func main() {
	fmt.Println(postorder(
		&Node{
			Val: 1,
			Children: []*Node{
				{
					Val: 3,
					Children: []*Node{
						{
							Val: 5,
						},
						{
							Val: 6,
						},
					},
				},
				{
					Val: 2,
				},
				{
					Val: 4,
				},
			},
		},
	))
}
