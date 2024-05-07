package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	var result *ListNode
	var inverted *ListNode
	current := head
	remainder := 0

	for current != nil {
		if inverted == nil {
			inverted = &ListNode{
				Val: current.Val,
			}
		} else {
			tmp := &ListNode{
				Val: current.Val,
			}
			tmp.Next = inverted
			inverted = tmp
		}
		current = current.Next
	}

	for inverted != nil {
		val := remainder + inverted.Val*2

		if val > 9 {
			tmpVal := val
			val = tmpVal % 10
			remainder = tmpVal / 10
		} else {
			remainder = 0
		}

		if result == nil {
			result = &ListNode{
				Val: val,
			}
		} else {
			tmp := &ListNode{
				Val: val,
			}
			tmp.Next = result
			result = tmp
		}

		inverted = inverted.Next
	}

	if remainder > 0 {
		tmp := &ListNode{
			Val: remainder,
		}

		tmp.Next = result
		result = tmp
	}

	return result
}

func printResult(node *ListNode) {
	current := node

	for current != nil {
		fmt.Print(current.Val)
		current = current.Next
	}

	fmt.Println()
}

func main() {
	printResult(doubleIt(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 9,
			},
		},
	}))

	printResult(doubleIt(&ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}))
}
