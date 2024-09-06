package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	hm := make(map[int]bool)
	var modified *ListNode
	current := modified

	for _, num := range nums {
		hm[num] = true
	}

	for head != nil {
		if !hm[head.Val] {
			if modified == nil {
				modified = &ListNode{
					Val: head.Val,
				}
				current = modified
			} else {
				node := &ListNode{
					Val: head.Val,
				}
				current.Next = node
				current = current.Next
			}
		}

		head = head.Next
	}

	return modified
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		print(current.Val)
		if current.Next != nil {
			print("->")
		}
		current = current.Next
	}
	println()
}

func main() {
	printList(modifiedList([]int{1, 2, 3}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}))
}
