package main

import (
	"fmt"
)

func finalPrices(prices []int) []int {
	sortedPairs := make([]int, len(prices)-1)

	for idx, price := range prices {
		if idx == 0 {
			continue
		}
		sortedPairs[idx-1] = price
	}

	result := make([]int, len(prices))

	for idx, price := range prices {
		i, found := 0, false

		for i < len(sortedPairs) {
			if sortedPairs[i] <= price {
				found = true
				result[idx] = price - sortedPairs[i]
				break
			}
			i++
		}

		if !found {
			result[idx] = price
		}

		if len(sortedPairs) >= 1 {
			sortedPairs = sortedPairs[1:]
		}
	}

	return result
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
}
