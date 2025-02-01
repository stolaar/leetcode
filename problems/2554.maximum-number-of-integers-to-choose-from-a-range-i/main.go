package main

func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := map[int]bool{}

	for _, bannedNum := range banned {
		if bannedNum > maxSum {
			continue
		}

		bannedMap[bannedNum] = true
	}

	total, sum := 0, 0

	for i := 1; i <= n; i++ {
		if sum == maxSum {
			return total
		}

		if _, ok := bannedMap[i]; ok {
			continue
		}

		if sum+i > maxSum {
			return total
		}

		total += 1
		sum += i
	}

	return total
}

func main() {
	println(maxCount([]int{1, 6, 5}, 5, 6))
}
