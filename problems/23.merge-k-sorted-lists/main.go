package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (list *LinkedList) insertIntoList(val int) {
	newNode := &ListNode{
		Val: val,
	}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

type Item struct {
	value    int
	priority int
	index    int
}
type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(val any) {
	value := val.(int)
	*pq = append(*pq, value)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := PriorityQueue{}
	heap.Init(&pq)
	result := LinkedList{}

	for _, list := range lists {
		if list == nil {
			continue
		}
		current := list
		for current.Next != nil {
			heap.Push(&pq, current.Val)
			current = current.Next
		}
		if current != nil {
			heap.Push(&pq, current.Val)
		}
	}

	for pq.Len() > 0 {
		result.insertIntoList(heap.Pop(&pq).(int))
	}

	return result.head
}

func printList(list *ListNode) {
	for list != nil {
		print(list.Val)
		list = list.Next
	}
	println()
}

func intToNodeList(nums []int) *ListNode {
	var head *ListNode

	for _, num := range nums {
		newNode := &ListNode{
			Val: num,
		}
		if head == nil {
			head = newNode
			continue
		}
		current := head

		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	return head
}

func main() {
	printList(mergeKLists([]*ListNode{
		intToNodeList([]int{1, 4, 5}),
		intToNodeList([]int{2, 6}),
		intToNodeList([]int{1, 3, 4}),
	}))

	printList(mergeKLists([]*ListNode{}))
	node := &ListNode{}
	printList(mergeKLists([]*ListNode{node}))
}
