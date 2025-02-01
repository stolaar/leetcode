package main

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	current := result

	for list1 != nil || list2 != nil {
		if list1 == nil {
			node := &ListNode{
				Val: list2.Val,
			}

			if result == nil {
				result = &ListNode{
					Val: list2.Val,
				}
				current = result
			} else {
				current.Next = node
				current = current.Next
			}

			list2 = list2.Next
			continue
		}

		if list2 == nil {
			node := &ListNode{
				Val: list1.Val,
			}

			if result == nil {
				result = &ListNode{
					Val: list1.Val,
				}
				current = result
			} else {
				current.Next = node
				current = current.Next
			}
			list1 = list1.Next
		}

		next := min(list1.Val, list2.Val)

		node := &ListNode{
			Val: next,
		}
		if result == nil {
			result = &ListNode{
				Val: next,
			}
			current = result
		} else {
			current.Next = node
			current = current.Next
		}

		if next == list1.Val {
			list1 = list1.Next
			continue
		}

		if next == list2.Val {
			list2 = list2.Next
			continue
		}
	}

	return result
}

func main() {
}

