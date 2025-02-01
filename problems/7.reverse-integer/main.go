package main

import "math"

func reverse(x int) int {
	reversed, negative := 0, x < 0

	if negative {
		x = int(math.Abs(float64(x)))
	}

	for x > 0 {
		rm := x % 10
		reversed = (reversed * 10) + rm
		x = x / 10
	}

	if negative {
		reversed = reversed - (2 * reversed)
	}

	if reversed > math.MaxInt32 && !negative {
		return 0
	}

	if reversed < math.MinInt32 {
		return 0
	}

	return reversed
}

func main() {
	println(reverse(2147483649))
	println(reverse(-2331))
}
