package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nm, longest := make(map[int]bool), 1

	for _, num := range nums {
		nm[num] = true
	}

	for k := range nm {
		sl, cl, cr := 1, k, k

		for {
			if _, ok := nm[cl-1]; ok {
				sl += 1
				delete(nm, cl-1)
				cl -= 1
			} else {
				break
			}
		}

		for {
			if _, ok := nm[cr+1]; ok {
				sl += 1
				delete(nm, cr+1)
				cr += 1
			} else {
				break
			}
		}

		if sl > longest {
			longest = sl
		}
	}

	return longest
}

func main() {
	println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
