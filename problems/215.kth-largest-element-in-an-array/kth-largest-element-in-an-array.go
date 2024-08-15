package main

import "container/heap"

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
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

func findKthLargest(nums []int, k int) int {
	pq := PriorityQueue{}
	heap.Init(&pq)
	for _, num := range nums {
		heap.Push(&pq, num)
	}

	var result int

	for i := 0; i < k; i++ {
		result = heap.Pop(&pq).(int)
	}

	return result
}

func main() {
	println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	println(findKthLargest([]int{-1, 2, 0}, 2)) // expected 0 got -1
}
