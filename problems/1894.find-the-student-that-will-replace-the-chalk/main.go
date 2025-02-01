package main

func chalkReplacer(chalk []int, k int) int {
	current, i := k, 0

	for current >= 0 {
		if chalk[i] > current {
			break
		}

		current -= chalk[i]

		next := i + 1
		if next > len(chalk)-1 {
			i = 0
			continue
		}

		i = next
	}

	return i
}

func main() {
	println(chalkReplacer([]int{5, 1, 5}, 22))
	println(chalkReplacer([]int{3, 4, 1, 2}, 25))
}
