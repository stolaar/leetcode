package main

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
type ListNode struct {
    Val int
    Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode

	var current *ListNode

	var inverted *ListNode

	var carry int = 0

	var currentL1 = l1
	var currentL2 = l2

	for currentL1 != nil || currentL2 != nil {
		val := carry

		if currentL1 != nil {
			val += currentL1.Val
		}

		if currentL2 != nil {
			val += currentL2.Val
		}

		if val > 9 {
			slice := []int{}

			for val > 0 {
				slice = append(slice, val%10)
				val /= 10
			}
			carry = slice[1]
			val = slice[0]
		} else {
			carry = 0
		}

		if currentL1 != nil {
			currentL1 = currentL1.Next
		}

		if currentL2 != nil {
			currentL2 = currentL2.Next
		}

		if result == nil {
			result = &ListNode{
				Val: val,
			}
		} else {
			current = &ListNode{
				Val: val,
			}
			current.Next = result
			result = current
		}
	}

	if carry > 0 {
		current = &ListNode{
			Val: carry,
		}
		current.Next = result
		result = current
	}

	for result != nil {
		if inverted == nil {
			inverted = &ListNode{
				Val: result.Val,
			}
		} else {
			current = &ListNode{
				Val: result.Val,
			}
			current.Next = inverted
			inverted = current
		}
		result = result.Next
	}

	return inverted
}

// @lc code=end
