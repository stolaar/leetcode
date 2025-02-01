package main

import "strconv"

func getLucky(s string, k int) int {
	ns, result := "", 0

	for _, r := range s {
		ns += strconv.Itoa(int(r - 'a' + 1))
	}

	for i := 0; i < k; i++ {
		c := 0
		for _, r := range ns {
			c += int(r - '0')
		}
		ns = strconv.Itoa(c)
		if i == k-1 {
			result = c
		}
	}

	return result
}

func main() {
	println(getLucky("iiii", 1))
	println(getLucky("leetcode", 2))
	println(getLucky("zbax", 2))
}
