package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func lexicalOrder(n int) []int {
	s := make([]int, n)

	for i := 1; i <= n; i++ {
		s[i-1] = i
	}

	slices.SortFunc(s, func(a, b int) int {
		as, bs := strconv.Itoa(a), strconv.Itoa(b)

		return strings.Compare(as, bs)
	})

	return s
}

func main() {
	fmt.Println(lexicalOrder(13))
}
