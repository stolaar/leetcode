package main

func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	if n <= 3 {
		return n
	}

	stepsCount, clipboard, left := 2, 1, n-2

	for left > 0 {
		printed := n - left

		if left%printed == 0 && printed < left {
			stepsCount += 2
			clipboard = printed
			left -= printed
			continue
		}

		stepsCount += 1
		left -= clipboard
	}

	return stepsCount
}

func main() {
	println(minSteps(3))
	println(minSteps(10))
}
