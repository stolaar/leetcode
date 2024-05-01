package main

import "fmt"

/*
 * @lc app=leetcode id=1137 lang=golang
 *
 * [1137] N-th Tribonacci Number
 */

// @lc code=start
func tribonacci(n int) int {
	if n <= 2 {
		return n
	}
	a := 0
	b := 1
	c := 1
	var d int

	for i := 2; i < n; i++ {
		d = a + b + c
		a = b
		b = c
		c = d
	}

	return d
}

// @lc code=end

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
