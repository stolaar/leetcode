package main

import (
	"fmt"
	"sort"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	frMap := make(map[float32][]int)
	frArr := []float32{}

	for i, val := range arr {
		for j := len(arr) - 1; j > i; j-- {
			fr := float32(val) / float32(arr[j])
			frMap[fr] = []int{val, arr[j]}
			frArr = append(frArr, fr)
		}
	}

	sort.Slice(frArr, func(i int, j int) bool {
		return frArr[i] < frArr[j]
	})

	return frMap[frArr[k-1]]
}

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))
}
