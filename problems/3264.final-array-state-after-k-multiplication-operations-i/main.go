package main

import (
	"container/heap"
	"fmt"
)

type PQ [][2]int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i][0] == pq[j][0] {
		return pq[i][1] < pq[j][1]
	}
	return pq[i][0] < pq[j][0]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(val any) {
	value := val.([2]int)
	*pq = append(*pq, value)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func getFinalState(nums []int, k int, multiplier int) []int {
	pq := &PQ{}

	heap.Init(pq)

	for idx, num := range nums {
		heap.Push(pq, [2]int{num, idx})
	}

	for i := 0; i < k; i++ {
		curr := heap.Pop(pq).([2]int)

		updated := curr[0] * multiplier

		nums[curr[1]] = updated
		heap.Push(pq, [2]int{updated, curr[1]})
		fmt.Println(nums)
	}

	return nums
}

func main() {
	fmt.Println(getFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
}
