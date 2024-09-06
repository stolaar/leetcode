package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	vm := make(map[int]int)
	var result *ListNode
	currentResult := result

	current := head
	for current != nil {
		vm[current.Val] += 1

		if vm[current.Val] <= 1 {
			if result == nil {
				result = &ListNode{
					Val: current.Val,
				}
				currentResult = result
				continue
			}
			node := &ListNode{
				Val: current.Val,
			}
			currentResult.Next = node
			currentResult = currentResult.Next
		}
		current = current.Next
	}

	return result
}

func main() {
}
