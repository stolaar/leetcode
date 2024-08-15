package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	c := a / b
	r := a - (b * c)

	return gcd(b, r)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	current := head
	printList(head)

	for current.Next != nil {
		node := &ListNode{
			Val:  gcd(current.Val, current.Next.Val),
			Next: current.Next,
		}

		current.Next = node
		current = node.Next
	}

	return head
}

func printList(head *ListNode) {
	current := head

	print("list: ")
	for current != nil {
		print(current.Val)
		if current.Next != nil {
			print(" -> ")
		}
		current = current.Next
	}
	println("")
}

func createList(nums []int) *ListNode {
	result := &ListNode{
		Val: nums[0],
	}

	current := result

	for _, num := range nums[1:] {
		tmp := &ListNode{
			Val: num,
		}
		current.Next = tmp

		current = current.Next
	}

	return result
}

func main() {
	printList(insertGreatestCommonDivisors(createList([]int{18, 6, 10, 3})))
}
