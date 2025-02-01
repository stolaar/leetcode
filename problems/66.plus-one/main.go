package main

import "fmt"

func plusOne(digits []int) []int {
	c := 1

	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i]

		num += c
		c = 0

		if num >= 10 {
			digits[i] = 0
			c = 1
			continue
		}

		digits[i] += 1
		break
	}

	if c > 0 {
		digits = append([]int{c}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	fmt.Println(plusOne([]int{1, 2}))
}
