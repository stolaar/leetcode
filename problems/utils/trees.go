package utils

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// Initialize the head of the linked list
	head := &ListNode{Val: nums[0]}
	current := head

	// Iterate over the rest of the elements and create nodes
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func BuildTree(data string) *TreeNode {
	if data == "" {
		return nil
	}

	values := strings.Split(data, ",")

	if values[0] == "null" {
		return nil
	}

	rootVal, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(values) {
		node := queue[0]
		queue = queue[1:]

		if values[index] != "null" {
			leftVal, _ := strconv.Atoi(values[index])
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}
		index++

		if index < len(values) && values[index] != "null" {
			rightVal, _ := strconv.Atoi(values[index])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

func PrintList(head *ListNode) {
	current, count := head, 0

	for current != nil {
		count++
		print(current.Val)
		if current.Next != nil {
			print("->")
		}
		current = current.Next
	}
	println()
	println("length", count)
}

func PrintTree(root *TreeNode, space int) {
	if root == nil {
		return
	}

	space += 7

	PrintTree(root.Right, space)

	println()

	for i := 10; i < space; i++ {
		print(" ")
	}

	println(root.Val)

	PrintTree(root.Left, space)
}
