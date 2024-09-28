package main

import "github.com/stolaar/leetcode/problems/utils"

type ListNode = utils.ListNode

type MyCircularDeque struct {
	maxSize int
	size    int
	list    *ListNode
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		maxSize: k,
		size:    0,
	}
}

func (deq *MyCircularDeque) InsertFront(value int) bool {
	if deq.size+1 > deq.maxSize {
		return false
	}
	node := &ListNode{
		Val: value,
	}

	if deq.list == nil {
		deq.list = node
		deq.size += 1
		return true
	}

	tmp := deq.list
	deq.list = node
	node.Next = tmp

	deq.size += 1

	return true
}

func (deq *MyCircularDeque) InsertLast(value int) bool {
	if deq.size+1 > deq.maxSize {
		return false
	}
	node := &ListNode{
		Val: value,
	}

	if deq.list == nil {
		deq.list = node
		deq.size += 1
		return true
	}

	current := deq.list
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	deq.size += 1
	return true
}

func (deq *MyCircularDeque) DeleteFront() bool {
	if deq.size-1 < 0 {
		return false
	}

	deq.size -= 1
	deq.list = deq.list.Next
	return true
}

func (deq *MyCircularDeque) DeleteLast() bool {
	if deq.size-1 < 0 {
		return false
	}

	current, next := deq.list, deq.list.Next

	if next == nil {
		deq.list = nil
		deq.size -= 1
		return true
	}
	for next != nil {
		if next.Next == nil {
			current.Next = nil
		}

		current = current.Next
		next = next.Next
	}

	deq.size -= 1

	return true
}

func (deq *MyCircularDeque) GetFront() int {
	if deq.list == nil {
		return -1
	}
	return deq.list.Val
}

func (deq *MyCircularDeque) GetRear() int {
	if deq.list == nil {
		return -1
	}
	current := deq.list

	for current.Next != nil {
		current = current.Next
	}

	return current.Val
}

func (deq *MyCircularDeque) IsEmpty() bool {
	return deq.size == 0
}

func (deq *MyCircularDeque) IsFull() bool {
	return deq.size == deq.maxSize
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {
	obj := Constructor(3)
	println(obj.InsertFront(9))
	println(obj.DeleteLast())
	println(obj.GetFront())
	println(obj.GetRear())
	println(obj.InsertLast(1))
	println(obj.InsertLast(2))
	println(obj.InsertFront(3))
	println(obj.InsertFront(4))
	println(obj.GetRear())
	println(obj.IsFull())
	println(obj.DeleteLast())
	println(obj.InsertFront(4))
	println(obj.GetFront())
}
