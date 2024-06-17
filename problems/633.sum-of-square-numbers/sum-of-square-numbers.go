package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	result := false
	s, n := float64(0), math.Ceil(math.Sqrt(float64(c)))

	for s <= n {
		sum := int(math.Pow(s, float64(2)) + math.Pow(n, float64(2)))

		if sum == c {
			result = true
			break
		}

		if sum > c {
			n--
			continue
		}

		if sum < c {
			s++
		}
	}

	return result
}

func printSolution(c int, expected bool) {
	fmt.Println(c, judgeSquareSum(c), " - expected", expected)
}

func main() {
	printSolution(5, true)
	printSolution(3, false)
	printSolution(8, false)
	printSolution(53, true)
	printSolution(4, true)
	printSolution(2, true)
	printSolution(1, true)
	printSolution(0, true)
}
