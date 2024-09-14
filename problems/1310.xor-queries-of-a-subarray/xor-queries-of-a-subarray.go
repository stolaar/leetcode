package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	ps, result := make([]int, len(arr)), []int{}
	i, j := 0, 1

	ps[0] = arr[0]
	ps[len(arr)-1] = arr[len(arr)-1]

	for i < len(arr)-1 {
		ps[i] = arr[i] ^ arr[j]

		i += 1
		if j < len(arr)-1 {
			j += 1
		}
	}

	for _, query := range queries {
		left, right := query[0], query[1]
		val := ps[left]

		if left == right {
			val = arr[left]
			result = append(result, val)
			continue
		}

		diff := right - left
		if diff == 1 {
			val = ps[left]
			result = append(result, val)
			continue
		}

		if (diff+1)%2 != 0 {

			for i := left + 2; i < right; i += 2 {
				val ^= ps[i]
			}

			val ^= arr[right]
			result = append(result, val)
			continue
		}

		for i := left + 2; i <= right; i += 2 {
			val ^= ps[i]
		}

		result = append(result, val)
	}

	return result
}

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
	fmt.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
	fmt.Println(xorQueries([]int{1, 15, 8, 4}, [][]int{{2, 2}, {0, 0}, {0, 3}, {1, 2}}))
	fmt.Println(xorQueries([]int{1, 11, 14, 10, 7}, [][]int{{0, 2}, {0, 0}, {2, 3}, {1, 4}}))
}
